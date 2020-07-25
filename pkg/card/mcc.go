package card

func TranslateMCC(code string) string {

	mcc := map[string]string{
		"1111": "Супермаркеты",
		"2222": "Рестораны",
		"5555": "Заправки",
	}

	result := "Категория не указана"

	for i := range mcc {
		if i == code {
			result = mcc[i]
		}
	}
	return result
}
