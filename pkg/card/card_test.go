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
				transaction: MakeTransactions(),
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
				transactions: MakeTransactions(),
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
				transactions: MakeTransactions(),
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
				transactions: MakeTransactions(),
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

func BenchmarkCategorization(b *testing.B) {
	transactions := MakeTransactions()
	want := map[string]int64{
		"Рестораны":    200,
		"Супермаркеты": 100,
	}
	b.ResetTimer() // сбрасываем таймер, т.к. сама генерация транзакций достаточно ресурсоёмка
	for i := 0; i < b.N; i++ {
		result := GetMap(transactions, 1)
		b.StopTimer() // останавливаем таймер, чтобы время сравнения не учитывалось
		if !reflect.DeepEqual(result, want) {
			b.Fatalf("invalid result, got %v, want %v", result, want)
		}
		b.StartTimer() // продолжаем работу таймера
	}
}

func BenchmarkCategorizationByMutex(b *testing.B) {
	transactions := MakeTransactions()
	want := map[string]int64{
		"Рестораны":    200,
		"Супермаркеты": 100,
	}
	b.ResetTimer() // сбрасываем таймер, т.к. сама генерация транзакций достаточно ресурсоёмка
	for i := 0; i < b.N; i++ {
		result := GetMapByMutex(transactions, 1, 1000)
		b.StopTimer() // останавливаем таймер, чтобы время сравнения не учитывалось
		if !reflect.DeepEqual(result, want) {
			b.Fatalf("invalid result, got %v, want %v", result, want)
		}
		b.StartTimer() // продолжаем работу таймера
	}
}

func BenchmarkCategorizationByChannel(b *testing.B) {
	transactions := MakeTransactions()
	want := map[string]int64{
		"Рестораны":    200,
		"Супермаркеты": 100,
	}
	b.ResetTimer() // сбрасываем таймер, т.к. сама генерация транзакций достаточно ресурсоёмка
	for i := 0; i < b.N; i++ {
		result := GetMapByChannel(transactions, 1, 1000)
		b.StopTimer() // останавливаем таймер, чтобы время сравнения не учитывалось
		if !reflect.DeepEqual(result, want) {
			b.Fatalf("invalid result, got %v, want %v", result, want)
		}
		b.StartTimer() // продолжаем работу таймера
	}
}

func BenchmarkCategorizationByMutex2(b *testing.B) {
	transactions := MakeTransactions()
	want := map[string]int64{
		"Рестораны":    200,
		"Супермаркеты": 100,
	}
	b.ResetTimer() // сбрасываем таймер, т.к. сама генерация транзакций достаточно ресурсоёмка
	for i := 0; i < b.N; i++ {
		result := GetMapByMutex2(transactions, 1, 1000)
		b.StopTimer() // останавливаем таймер, чтобы время сравнения не учитывалось
		if !reflect.DeepEqual(result, want) {
			b.Fatalf("invalid result, got %v, want %v", result, want)
		}
		b.StartTimer() // продолжаем работу таймера
	}
}
