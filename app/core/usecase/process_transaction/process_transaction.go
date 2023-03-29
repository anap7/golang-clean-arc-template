package process_transaction

import (
	"github.com/anap7/golang-clean-arc-template/app/core/entity"
	"github.com/anap7/golang-clean-arc-template/app/core/repository"
)

type ProcessTransaction struct {
	//Implementando a interface repository
	Repository repository.TransactionRepository
}

// Construtor - Objeto com o repository instancioado
func NewProcessTransaction(repository repository.TransactionRepository) *ProcessTransaction {
	return &ProcessTransaction{Repository: repository}
}

func (p *ProcessTransaction) Execute(input TransactionDtoInput) (TransactionDtoOutput, error) {
	transaction := entity.NewTransaction()
	transaction.ID = input.ID
	transaction.AccountID = input.AccountID
	transaction.Amount = input.Amount
	invalidTransaction := transaction.IsValid()

	if invalidTransaction != nil {
		err := p.Repository.Insert(transaction.ID, transaction.AccountID, transaction.Amount, "rejected", invalidTransaction.Error())
		if err != nil {
			return TransactionDtoOutput{}, err
		}

		output := TransactionDtoOutput{
			ID:           transaction.ID,
			Status:       "rejected",
			ErrorMessage: invalidTransaction.Error(),
		}

		return output, nil
	}

	err := p.Repository.Insert(transaction.ID, transaction.AccountID, transaction.Amount, "approved", "")
	if err != nil {
		return TransactionDtoOutput{}, err
	}

	output := TransactionDtoOutput{
		ID:           transaction.ID,
		Status:       "approved",
		ErrorMessage: invalidTransaction.Error(),
	}

	return output, nil
}
