package card

import (
	"reflect"
	"testing"
)

func TestGetMap(t *testing.T) {
	type args struct {
		transaction []Transaction
		id          int64
	}
	tests := []struct {
		name string
		args args
		want map[string]int64
	}{
		{
			name: "1",
			args: args{
				transaction: MakeTransactions1(),
				id:          1,
			},
			want: map[string]int64{
				"Рестораны":    200,
				"Супермаркеты": 100,
			},
		},
	}
	for _, tt := range tests {
		if got := GetMap(tt.args.transaction, tt.args.id); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("GetMap() = %v, want %v", got, tt.want)
		}
	}
}

func TestGetMapByMutex(t *testing.T) {
	type args struct {
		transactions []Transaction
		id           int64
		goroutines   int64
	}
	tests := []struct {
		name string
		args args
		want map[string]int64
	}{
		{
			name: "1",
			args: args{
				transactions: MakeTransactions1(),
				id:           1,
				goroutines:   10000,
			},
			want: map[string]int64{
				"Рестораны":    200,
				"Супермаркеты": 100,
			},
		},
	}
	for _, tt := range tests {
		if got := GetMapByMutex(tt.args.transactions, tt.args.id, tt.args.goroutines); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("GetMapByMutex() = %v, want %v", got, tt.want)
		}
	}
}

func TestGetMapByChannel(t *testing.T) {
	type args struct {
		transactions []Transaction
		id           int64
		partCount    int64
	}
	tests := []struct {
		name string
		args args
		want map[string]int64
	}{
		{
			name: "1",
			args: args{
				transactions: MakeTransactions1(),
				id:           1,
				partCount:    10000,
			},
			want: map[string]int64{
				"Рестораны":    200,
				"Супермаркеты": 100,
			},
		},
	}
	for _, tt := range tests {
		if got := GetMapByChannel(tt.args.transactions, tt.args.id, tt.args.partCount); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("GetMapByChannel() = %v, want %v", got, tt.want)
		}
	}
}

func TestGetMapByMutex2(t *testing.T) {
	type args struct {
		transactions []Transaction
		id           int64
		goroutines   int64
	}
	tests := []struct {
		name string
		args args
		want map[string]int64
	}{
		{
			name: "1",
			args: args{
				transactions: MakeTransactions1(),
				id:           1,
				goroutines:   10000,
			},
			want: map[string]int64{
				"Рестораны":    200,
				"Супермаркеты": 100,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMapByMutex2(tt.args.transactions, tt.args.id, tt.args.goroutines); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMapByMutex2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func MakeTransactions1() []Transaction {
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
