package card

import (
	"sync"
)

// variant 1
func GetMap(transaction []Transaction, id int64) map[string]int64 {
	total := make(map[string]int64)
	for i := range transaction {
		if transaction[i].Id == id {
			category := TranslateMCC(transaction[i].Mcc)
			total[category] += transaction[i].Amount
		}
	}
	return total
}

// variant 2
func GetMapByMutex(transactions []Transaction, id, goroutines int64) map[string]int64 {
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	result := make(map[string]int64)

	for i := int64(0); i < goroutines; i++ {
		wg.Add(1)
		partSize := int64(len(transactions)) / goroutines
		part := transactions[i*partSize : (i+1)*partSize]
		go func() {
			m := GetMap(part, 1)

			mu.Lock()
			for k, value := range m {
				switch k {
				case "Рестораны":
					result["Рестораны"] += value
				case "Супермаркеты":
					result["Супермаркеты"] += value
				default:
					result["Остальное"] += value
				}
			}
			mu.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()

	return result
}

// variant 3
func GetMapByChannel(transactions []Transaction, id, partCount int64) map[string]int64 {
	result := make(map[string]int64)
	ch := make(chan map[string]int64)

	for i := int64(0); i < partCount; i++ {
		partSize := int64(len(transactions)) / partCount
		part := transactions[i*partSize : (i+1)*partSize]
		go func(ch chan<- map[string]int64) {
			ch <- GetMap(part, 1)
		}(ch)
	}

	finished := int64(0)
	for valueCh := range ch {
		for k, value := range valueCh {
			switch k {
			case "Рестораны":
				result["Рестораны"] += value
			case "Супермаркеты":
				result["Супермаркеты"] += value
			default:
				result["Остальное"] += value
			}
		}
		finished++
		if finished == partCount {
			break
		}
	}

	return result
}

// variant 4
func GetMapByMutex2(transactions []Transaction, id, goroutines int64) map[string]int64 {
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}
	result := make(map[string]int64)

	for i := int64(0); i < goroutines; i++ {
		wg.Add(1)
		partSize := int64(len(transactions)) / goroutines
		part := transactions[i*partSize : (i+1)*partSize]
		go func() {
			for _, i := range part {
				if i.Id == id {
					mu.Lock()
					category := TranslateMCC(i.Mcc)
					result[category] += i.Amount
					mu.Unlock()
				}
			}
			wg.Done()
		}()
	}
	wg.Wait()

	return result
}

// struct

type Card struct {
	Id           int64
	Issuer       string
	Balance      int64
	Currency     string
	Number       string
	Transactions []Transaction
}

type Transaction struct {
	Id     int64
	Amount int64
	Mcc    string
}

func MakeTransactions() []Transaction {
	const users = 10
	const transactionsPerUser = 10_00
	const transactionAmount = 1_00
	transactions := make([]Transaction, users*transactionsPerUser)
	for index := range transactions {
		switch index % 100 {
		case 0:
			transactions[index] = Transaction{Id: 1, Amount: 1, Mcc: "1111"} // Например, каждая 100-ая транзакция в банке от нашего юзера в категории такой-то
		case 20:
			transactions[index] = Transaction{Id: 1, Amount: 2, Mcc: "2222"} // Например, каждая 120-ая транзакция в банке от нашего юзера в категории такой-то
		default:
			transactions[index] = Transaction{Id: 2, Amount: transactionAmount, Mcc: "5555"} // Транзакции других юзеров, нужны для "общей" массы
		}
	}
	return transactions
}
