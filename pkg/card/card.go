package card

type CodeMCC struct {
	Id int64
}


type Transaction struct {
	Id int64
	AmountTransaction int64
	Date int64
	CodeMCC []*CodeMCC{}
	Status string
}

type Card struct {
	Id int64
	Issuer string
	Balance int64
	Currency string
	Number string
	Transaction []*Transaction{}
}

func AddTransaction(card *Card, transaction *Transaction) {
	card.Balance -= transaction.AmountTransaction
}

func SumByMCC (transaction []*Transaction, CodeMCC []*CodeMCC) int64 {
    var result int64
    count:= len(transaction)
    for i:=0; i < count; i++{
    	if MCCMatch(){
		}
    	result += transaction[i].AmountTransaction
	}
	return result
}
func MCCMatch () int64 {
	var result int64
	count:= len(CodeMCC)
	for j:= 0; j<count; j++{
		result = CodeMCC[j].Id
	}
	return result
}

func MCCMatch () {

}