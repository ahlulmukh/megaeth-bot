package main

import (
	"bepop-teko/internal/bepop"
	"bepop-teko/internal/pkg/ethclient"
	"bepop-teko/internal/pkg/logger"
	"bepop-teko/internal/teko"
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"math/big"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/fatih/color"
)

const (
	bepopAddress  = "0x4eb2bd7bee16f38b1f4a0a5796fffd028b6040e9"
	tkETH         = "0x176735870dc6c22b4ebfbf519de2ce758de78d94"
	tkUSDC        = "0xfaf334e157175ff676911adcf0964d7f54f2c424"
	tkBTC         = "0xf82ff0799448630eb56ce747db840a2e02cde4d8"
	cUSDC         = "0xe9b6e75c243b6100ffcb1c66e8f78f96feea727f"
	configFile    = "config.json"
	proxyFile     = "proxy.txt"
	minSwapAmount = 0.001
	maxSwapAmount = 0.1
)

type Config struct {
	PrivateKeys []string `json:"private_keys"`
}

func displayBanner() {
	color.Cyan(`
░█▄█░█▀▀░█▀▀░█▀█░█▀▀░▀█▀░█░█
░█░█░█▀▀░█░█░█▀█░█▀▀░░█░░█▀█
░▀░▀░▀▀▀░▀▀▀░▀░▀░▀▀▀░░▀░░▀░▀`)
	fmt.Println()
	color.Yellow("  By : El Puqus Airdrop")
	color.Magenta("   github.com/ahlulmukh")
	color.Red(" Use it at your own risk")
	fmt.Println()
}

func main() {
	displayBanner()
	log := logger.New()
	proxyList, err := loadProxies(proxyFile)
	if err != nil {
		log.Error(fmt.Sprintf("Failed to load proxies: %v", err))
		os.Exit(1)
	}

	if proxyList != nil {
		log.Info(fmt.Sprintf("Loaded %d proxies from %s", len(proxyList), proxyFile))
	} else {
		log.Info("No proxies found, will use direct connection")
	}

	privateKeys, err := loadPrivateKeys(configFile)
	if err != nil {
		log.Error(fmt.Sprintf("Failed to load private keys: %v", err))
		os.Exit(1)
	}

	mode := selectMode()
	switch mode {
	case 1:
		runSingleAccountMode(privateKeys, proxyList, log)
	case 2:
		runMultiAccountMode(privateKeys, proxyList, log)
	case 3:
		addNewPrivateKey(log)
	default:
		log.Warning("Invalid mode selected.")
	}
}

func loadProxies(filename string) ([]string, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, nil
		}
		return nil, err
	}

	lines := strings.Split(string(content), "\n")
	var proxies []string
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" {
			proxies = append(proxies, line)
		}
	}

	if len(proxies) == 0 {
		return nil, nil
	}

	return proxies, nil
}

func selectMode() int {
	color.Cyan("\nChoose Mode:")
	color.Yellow("1. Single Account")
	color.Yellow("2. Multi Account")
	color.Yellow("3. Add New Private Key")
	fmt.Print("\nChoose mode (1/2/3): ")
	var mode int
	fmt.Scanln(&mode)
	return mode
}

func addNewPrivateKey(log *logger.Logger) {
	for {
		fmt.Print("Enter new private key: ")
		var privateKey string
		fmt.Scanln(&privateKey)
		privateKey = strings.TrimSpace(privateKey)

		if privateKey == "" {
			log.Warning("Private key cannot be empty.")
			return
		}

		privateKey = remove0xPrefix(privateKey)
		key, err := crypto.HexToECDSA(privateKey)
		if err != nil {
			log.Error(fmt.Sprintf("Invalid private key: %v", err))
			continue
		}

		existingKeys, err := loadPrivateKeys(configFile)
		if err != nil {
			log.Error(fmt.Sprintf("Failed to load private keys: %v", err))
			return
		}

		duplicate := false
		for _, existingKey := range existingKeys {
			if existingKey.D.Cmp(key.D) == 0 {
				log.Warning("Private key already exists in the configuration.")
				duplicate = true
				break
			}
		}
		if duplicate {
			continue
		}

		existingKeys = append(existingKeys, key)

		err = savePrivateKeys(configFile, existingKeys)
		if err != nil {
			log.Error(fmt.Sprintf("Failed to save private keys: %v", err))
			return
		}

		log.Success("New private key added successfully!")
		fmt.Print("Wanna add again? (y/n): ")
		var choice string
		fmt.Scanln(&choice)
		if strings.ToLower(choice) != "y" {
			break
		}
	}
}

func runSingleAccountMode(keys []*ecdsa.PrivateKey, proxyList []string, log *logger.Logger) {
	selectedIndex := selectAccount(keys, log)
	privateKey := keys[selectedIndex]
	var proxy string
	if proxyList != nil && len(proxyList) > 0 {
		proxy = proxyList[selectedIndex%len(proxyList)]
		log.Info(fmt.Sprintf("Using proxy: %s", proxy))
	}

	client, err := ethclient.New(context.Background(), "https://carrot.megaeth.com/rpc", proxy)
	if err != nil {
		log.Error(fmt.Sprintf("Failed to initialize Ethereum client: %v", err))
		return
	}
	defer client.Close()

	chainID, err := client.ChainID(context.Background())
	if err != nil {
		log.Error(fmt.Sprintf("Failed to fetch chain ID: %v", err))
		return
	}

	bepopSvc, err := bepop.NewService(client, common.HexToAddress(bepopAddress), chainID)
	if err != nil {
		log.Error(fmt.Sprintf("Failed to initialize Bepop service: %v", err))
		return
	}

	tekoSvc, err := teko.NewService(client, common.HexToAddress(tkETH), chainID)
	if err != nil {
		log.Error(fmt.Sprintf("Failed to initialize Teko service: %v", err))
		return
	}

	address := crypto.PubkeyToAddress(privateKey.PublicKey)
	ctx := context.Background()
	ethBalance, _ := client.BalanceAt(ctx, address, nil)
	wethBalance, _ := bepopSvc.BalanceOf(ctx, address)

	color.Green("\nBalance %s:", address.Hex())
	fmt.Printf("ETH:  %s\n", weiToEther(ethBalance))
	fmt.Printf("WETH: %s\n", weiToEther(wethBalance))

	for {
		displayMenu()
		fmt.Print("\nChoose menu: ")
		var choice int
		fmt.Scanln(&choice)

		if choice == 0 {
			break
		}

		processSingleAccountMenu(choice, privateKey, bepopSvc, tekoSvc, ethBalance, wethBalance, log, client)
	}
}

func selectAccount(keys []*ecdsa.PrivateKey, log *logger.Logger) int {
	color.Cyan("\nAccount Ready:")
	for i, key := range keys {
		addr := crypto.PubkeyToAddress(key.PublicKey).Hex()
		fmt.Printf("%d. %s\n", i+1, addr)
	}

	fmt.Print("\nChoose account (1-", len(keys), "): ")
	var choice int
	_, err := fmt.Scanln(&choice)
	if err != nil || choice < 1 || choice > len(keys) {
		log.Error("Not valid")
		os.Exit(1)
	}
	return choice - 1
}

func processSingleAccountMenu(choice int, privateKey *ecdsa.PrivateKey, bepopSvc *bepop.Service, tekoSvc *teko.Service, ethBalance, wethBalance *big.Int, log *logger.Logger, client *ethclient.Client) {

	switch choice {
	case 1:
		fmt.Print("Enter amount ETH to deposit (example: 0.1): ")
		var amountStr string
		fmt.Scanln(&amountStr)
		amount := etherToWei(amountStr)
		tx, err := bepopSvc.DepositETH(privateKey, amount)
		if err != nil {
			log.Error(fmt.Sprintf("Deposit failed: %v", err))
			return
		}
		log.Success(fmt.Sprintf("Deposit tx sent: %s", tx.Hash().Hex()))
		log.Success(fmt.Sprintf("Block explorer: https://megaexplorer.xyz/tx/%s", tx.Hash().Hex()))

	case 2:
		fmt.Print("Enter amount WETH to withdraw (example: 0.1): ")
		var amountStr string
		fmt.Scanln(&amountStr)
		amount := etherToWei(amountStr)
		tx, err := bepopSvc.Withdraw(privateKey, amount)
		if err != nil {
			log.Error(fmt.Sprintf("Withdraw failed: %v", err))
			return
		}
		log.Success(fmt.Sprintf("Withdraw tx sent: %s", tx.Hash().Hex()))
		log.Success(fmt.Sprintf("Block explorer: https://megaexplorer.xyz/tx/%s", tx.Hash().Hex()))

	case 3:
		fmt.Print("Enter max amount to swap (example: 0.5 ETH): ")
		var maxStr string
		fmt.Scanln(&maxStr)
		maxAmount := etherToWei(maxStr)

		fmt.Print("Enter number of random swaps to perform: ")
		var count int
		fmt.Scanln(&count)

		for i := 0; i < count; i++ {
			log.Info(fmt.Sprintf("Performing random swap %d/%d", i+1, count))

			swapETHorWETH := rand.Intn(2)
			amount := randomAmount(maxAmount)

			if swapETHorWETH == 0 {
				if ethBalance.Cmp(amount) < 0 {
					log.Warning("Not enough ETH balance to swap.")
					continue
				}
				log.Info(fmt.Sprintf("Swapping %s ETH -> WETH", weiToEther(amount)))
				tx, err := bepopSvc.DepositETH(privateKey, amount)
				if err != nil {
					log.Error(fmt.Sprintf("Random Deposit failed: %v", err))
					continue
				}
				log.Success(fmt.Sprintf("Swap successful! Amount: %s ETH -> WETH", weiToEther(amount)))
				log.Success(fmt.Sprintf("Transaction: https://megaexplorer.xyz/tx/%s", tx.Hash().Hex()))
			} else {
				if wethBalance.Cmp(amount) < 0 {
					log.Warning("Not enough WETH balance to swap.")
					continue
				}
				log.Info(fmt.Sprintf("Swapping %s WETH -> ETH", weiToEther(amount)))
				tx, err := bepopSvc.Withdraw(privateKey, amount)
				if err != nil {
					log.Error(fmt.Sprintf("Random Withdraw failed: %v", err))
					continue
				}
				log.Success(fmt.Sprintf("Swap successful! Amount: %s WETH -> ETH", weiToEther(amount)))
				log.Success(fmt.Sprintf("Transaction: https://megaexplorer.xyz/tx/%s", tx.Hash().Hex()))
			}

			if i < count-1 {
				delay := time.Duration(rand.Intn(3)+1) * time.Second
				log.Info(fmt.Sprintf("Waiting %s before next swap...", delay))
				time.Sleep(delay)
			}
		}

	case 4, 5, 6, 7, 8:
		processTekoMenu(choice, privateKey, tekoSvc, log)
	default:
		log.Warning("Menu tidak valid")
	}
}

func runMultiAccountMode(keys []*ecdsa.PrivateKey, proxyList []string, log *logger.Logger) {
	for {
		displayMultiAccountMenu()
		fmt.Print("\nChoose menu for all accounts: ")
		var choice int
		fmt.Scanln(&choice)

		if choice == 0 {
			break
		}

		switch choice {
		case 1, 2, 3:
			processMultiBepop(choice, keys, proxyList, log)
		case 4, 5, 6, 7, 8:
			processMultiTeko(choice, keys, proxyList, log)
		default:
			log.Warning("Invalid menu selected.")
		}
	}
}

func displayMultiAccountMenu() {
	color.Cyan("\n---- Menu Multi Akun ----")
	color.Yellow("1. ETH -> WETH (random amount)")
	color.Yellow("2. WETH -> ETH (random amount)")
	color.Yellow("3. Random swap ETH <-> WETH")
	color.Magenta("\n---- Teko Mint ----")
	color.Yellow("4. Mint tkETH")
	color.Yellow("5. Mint tkUSDC")
	color.Yellow("6. Mint tkBTC")
	color.Yellow("7. Mint cUSDC")
	color.Yellow("8. Mint All")
	color.Red("\n0. Back to Main Menu\n")
}

func processMultiBepop(choice int, keys []*ecdsa.PrivateKey, proxyList []string, log *logger.Logger) {
	for i, privateKey := range keys {
		var proxy string
		if proxyList != nil && len(proxyList) > 0 {
			proxy = proxyList[i%len(proxyList)]
			log.Info(fmt.Sprintf("Account %d using proxy: %s", i+1, proxy))
		}

		client, err := ethclient.New(context.Background(), "https://carrot.megaeth.com/rpc", proxy)
		if err != nil {
			log.Error(fmt.Sprintf("Account %d: Failed to initialize Ethereum client: %v", i+1, err))
			continue
		}
		defer client.Close()

		chainID, err := client.ChainID(context.Background())
		if err != nil {
			log.Error(fmt.Sprintf("Account %d: Failed to fetch chain ID: %v", i+1, err))
			continue
		}

		bepopSvc, err := bepop.NewService(client, common.HexToAddress(bepopAddress), chainID)
		if err != nil {
			log.Error(fmt.Sprintf("Account %d: Failed to initialize Bepop service: %v", i+1, err))
			continue
		}

		address := crypto.PubkeyToAddress(privateKey.PublicKey)
		log.Process(i+1, len(keys), fmt.Sprintf("Processing %s", address.Hex()))

		ctx := context.Background()
		ethBalance, _ := client.BalanceAt(ctx, address, nil)
		wethBalance, _ := bepopSvc.BalanceOf(ctx, address)

		minAmount := etherToWei(fmt.Sprintf("%f", minSwapAmount))
		maxAmount := etherToWei(fmt.Sprintf("%f", maxSwapAmount))
		amount := randomAmountInRange(minAmount, maxAmount)

		var txHash string
		var action string

		switch choice {
		case 1:
			if ethBalance.Cmp(amount) < 0 {
				log.Warning(fmt.Sprintf("%s: Insufficient ETH balance (need %s, have %s)",
					address.Hex(), weiToEther(amount), weiToEther(ethBalance)))
				continue
			}
			tx, err := bepopSvc.DepositETH(privateKey, amount)
			if err != nil {
				log.Error(fmt.Sprintf("%s: Deposit failed: %v", address.Hex(), err))
				continue
			}
			txHash = tx.Hash().Hex()
			action = fmt.Sprintf("%s ETH -> WETH", weiToEther(amount))

		case 2:
			if wethBalance.Cmp(amount) < 0 {
				log.Warning(fmt.Sprintf("%s: Insufficient WETH balance (need %s, have %s)",
					address.Hex(), weiToEther(amount), weiToEther(wethBalance)))
				continue
			}
			tx, err := bepopSvc.Withdraw(privateKey, amount)
			if err != nil {
				log.Error(fmt.Sprintf("%s: Withdraw failed: %v", address.Hex(), err))
				continue
			}
			txHash = tx.Hash().Hex()
			action = fmt.Sprintf("%s WETH -> ETH", weiToEther(amount))

		case 3:
			if rand.Intn(2) == 0 && ethBalance.Cmp(amount) >= 0 {
				tx, err := bepopSvc.DepositETH(privateKey, amount)
				if err != nil {
					log.Error(fmt.Sprintf("%s: Deposit failed: %v", address.Hex(), err))
					continue
				}
				txHash = tx.Hash().Hex()
				action = fmt.Sprintf("%s ETH -> WETH", weiToEther(amount))
			} else if wethBalance.Cmp(amount) >= 0 {
				tx, err := bepopSvc.Withdraw(privateKey, amount)
				if err != nil {
					log.Error(fmt.Sprintf("%s: Withdraw failed: %v", address.Hex(), err))
					continue
				}
				txHash = tx.Hash().Hex()
				action = fmt.Sprintf("%s WETH -> ETH", weiToEther(amount))
			} else {
				log.Warning(fmt.Sprintf("%s: Insufficient balance for swap", address.Hex()))
				continue
			}
		}

		log.Success(fmt.Sprintf("%s: %s", address.Hex(), action))
		log.Success(fmt.Sprintf("Tx: https://megaexplorer.xyz/tx/%s", txHash))

		if i < len(keys)-1 {
			time.Sleep(2 * time.Second)
		}
	}
}

func randomAmountInRange(min, max *big.Int) *big.Int {
	if min == nil || max == nil || min.Cmp(max) > 0 {
		return min
	}

	rand.Seed(time.Now().UnixNano())
	diff := new(big.Int).Sub(max, min)
	r := new(big.Int).Rand(rand.New(rand.NewSource(time.Now().UnixNano())), diff)
	return r.Add(r, min)
}

func processMultiTeko(choice int, keys []*ecdsa.PrivateKey, proxyList []string, log *logger.Logger) {
	for i, privateKey := range keys {
		var proxy string
		if proxyList != nil && len(proxyList) > 0 {
			proxy = proxyList[i%len(proxyList)]
			log.Info(fmt.Sprintf("Account %d using proxy: %s", i+1, proxy))
		}

		client, err := ethclient.New(context.Background(), "https://carrot.megaeth.com/rpc", proxy)
		if err != nil {
			log.Error(fmt.Sprintf("Account %d: Failed to initialize Ethereum client: %v", i+1, err))
			continue
		}
		defer client.Close()

		chainID, err := client.ChainID(context.Background())
		if err != nil {
			log.Error(fmt.Sprintf("Account %d: Failed to fetch chain ID: %v", i+1, err))
			continue
		}

		tekoSvc, err := teko.NewService(client, common.HexToAddress(tkETH), chainID)
		if err != nil {
			log.Error(fmt.Sprintf("Account %d: Failed to initialize Teko service: %v", i+1, err))
			continue
		}

		address := crypto.PubkeyToAddress(privateKey.PublicKey)
		log.Process(i+1, len(keys), fmt.Sprintf("Processing %s", address.Hex()))

		switch choice {
		case 4:
			mintToken(tekoSvc, privateKey, tkETH, "1", 18, log)
		case 5:
			mintToken(tekoSvc, privateKey, tkUSDC, "2000", 6, log)
		case 6:
			mintToken(tekoSvc, privateKey, tkBTC, "0.02", 8, log)
		case 7:
			mintToken(tekoSvc, privateKey, cUSDC, "1000", 6, log)
		case 8:
			mintToken(tekoSvc, privateKey, tkETH, "1", 18, log)
			time.Sleep(1 * time.Second)
			mintToken(tekoSvc, privateKey, tkUSDC, "2000", 6, log)
			time.Sleep(1 * time.Second)
			mintToken(tekoSvc, privateKey, tkBTC, "0.02", 8, log)
			time.Sleep(1 * time.Second)
			mintToken(tekoSvc, privateKey, cUSDC, "1000", 6, log)
		}

		if i < len(keys)-1 {
			time.Sleep(2 * time.Second)
		}
	}
}

func processTekoMenu(choice int, privateKey *ecdsa.PrivateKey, tekoSvc *teko.Service, log *logger.Logger) {
	switch choice {
	case 4:
		mintToken(tekoSvc, privateKey, tkETH, "1", 18, log)
	case 5:
		mintToken(tekoSvc, privateKey, tkUSDC, "2000", 6, log)
	case 6:
		mintToken(tekoSvc, privateKey, tkBTC, "0.02", 8, log)
	case 7:
		mintToken(tekoSvc, privateKey, cUSDC, "1000", 6, log)
	case 8:
		mintToken(tekoSvc, privateKey, tkETH, "1", 18, log)
		time.Sleep(1 * time.Second)
		mintToken(tekoSvc, privateKey, tkUSDC, "2000", 6, log)
		time.Sleep(1 * time.Second)
		mintToken(tekoSvc, privateKey, tkBTC, "0.02", 8, log)
		time.Sleep(1 * time.Second)
		mintToken(tekoSvc, privateKey, cUSDC, "1000", 6, log)
	default:
		log.Warning("Menu Teko tidak valid")
	}
}

func loadPrivateKeys(filename string) ([]*ecdsa.PrivateKey, error) {
	log := logger.New()
	_, err := os.Stat(filename)
	if err == nil {
		content, err := os.ReadFile(filename)
		if err != nil {
			return nil, err
		}
		var config Config
		err = json.Unmarshal(content, &config)
		if err != nil {
			return nil, err
		}

		var keys []*ecdsa.PrivateKey
		for _, privateKeyHex := range config.PrivateKeys {
			privateKeyHex = remove0xPrefix(privateKeyHex)

			privateKey, err := crypto.HexToECDSA(privateKeyHex)
			if err != nil {
				return nil, fmt.Errorf("invalid private key: %v", err)
			}
			keys = append(keys, privateKey)
		}
		return keys, nil
	} else if os.IsNotExist(err) {
		log.Warning("Config file not found. Please enter your private keys manually.")
		return promptPrivateKeys()
	} else {
		return nil, err
	}
}

func promptPrivateKeys() ([]*ecdsa.PrivateKey, error) {
	var privateKeys []*ecdsa.PrivateKey
	for {
		fmt.Print("Enter your private key: ")
		var privateKey string
		fmt.Scanln(&privateKey)
		privateKey = strings.TrimSpace(privateKey)

		if privateKey == "" {
			break
		}

		privateKey = remove0xPrefix(privateKey)

		key, err := crypto.HexToECDSA(privateKey)
		if err != nil {
			return nil, fmt.Errorf("invalid private key: %v", err)
		}

		privateKeys = append(privateKeys, key)

		fmt.Print("Do you want to add another private key? (y/n): ")
		var addMore string
		fmt.Scanln(&addMore)

		if strings.ToLower(addMore) != "y" {
			break
		}
	}

	err := savePrivateKeys(configFile, privateKeys)
	if err != nil {
		return nil, err
	}

	return privateKeys, nil
}

func remove0xPrefix(privateKey string) string {
	if strings.HasPrefix(privateKey, "0x") {
		return privateKey[2:]
	}
	return privateKey
}

func savePrivateKeys(filename string, keys []*ecdsa.PrivateKey) error {
	var privateKeyHexs []string
	for _, key := range keys {
		privateKeyHexs = append(privateKeyHexs, fmt.Sprintf("%x", key.D.Bytes()))
	}

	config := Config{PrivateKeys: privateKeyHexs}
	data, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}

	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func displayMenu() {
	color.Cyan("\n---- Bebop Exchange ----\n")
	color.Yellow("1. ETH -> WETH")
	color.Yellow("2. WETH -> ETH")
	color.Yellow("3. Random swap ETH <-> WETH")
	color.Magenta("\n---- Teko Mint ----")
	color.Yellow("4. Mint tkETH (1)")
	color.Yellow("5. Mint tkUSDC (2000)")
	color.Yellow("6. Mint tkBTC (0.02)")
	color.Yellow("7. Mint cUSDC (1000)")
	color.Green("8. Mint All")
	color.Red("\n0. Exit\n")
}

func amountToDecimals(amount string, decimals int) *big.Int {
	f, _ := new(big.Float).SetString(amount)
	multiplier := new(big.Float).SetFloat64(float64(1))
	multiplier.Quo(multiplier, big.NewFloat(1))
	for i := 0; i < decimals; i++ {
		multiplier.Mul(multiplier, big.NewFloat(10))
	}
	f.Mul(f, multiplier)
	result := new(big.Int)
	f.Int(result)
	return result
}

func mintToken(tekoSvc *teko.Service, privateKey *ecdsa.PrivateKey, tokenAddress string, amountStr string, decimals int, log *logger.Logger) {
	amount := amountToDecimals(amountStr, decimals)
	address := crypto.PubkeyToAddress(privateKey.PublicKey)
	tx, err := tekoSvc.Mint(privateKey, address, amount)
	if err != nil {
		log.Error(fmt.Sprintf("Mint failed: %v", err))
		return
	}
	log.Success("Mint succesffully")
	log.Success(fmt.Sprintf("Transaction : https://megaexplorer.xyz/tx/%s", tx.Hash().Hex()))
}

func weiToEther(wei *big.Int) string {
	f := new(big.Float).SetInt(wei)
	f.Quo(f, big.NewFloat(1e18))
	return f.Text('f', 6)
}

func etherToWei(amount string) *big.Int {
	f, _ := new(big.Float).SetString(amount)
	f.Mul(f, big.NewFloat(1e18))
	result := new(big.Int)
	f.Int(result)
	return result
}

func randomAmount(max *big.Int) *big.Int {
	rand.Seed(time.Now().UnixNano())
	r := new(big.Int).Rand(rand.New(rand.NewSource(time.Now().UnixNano())), max)
	return r
}
