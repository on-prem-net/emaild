package submissionbackend

import (
	"github.com/on-prem-net/emaild/dao"
	"github.com/on-prem-net/emaild/localdelivery"
	"github.com/asdine/genji"
)

type SubmissionBackend struct {
	accountsDAO         dao.AccountsDAO
	db                  *genji.DB
	localDelivery       *localdelivery.LocalDelivery
	mailboxesDAO        dao.MailboxesDAO
	mailboxMessagesDAO  dao.MailboxMessagesDAO
	messageRawBodiesDAO dao.MessageRawBodiesDAO
	messagesDAO         dao.MessagesDAO
}

func New(
	accountsDAO dao.AccountsDAO,
	db *genji.DB,
	localDelivery *localdelivery.LocalDelivery,
	mailboxesDAO dao.MailboxesDAO,
	mailboxMessagesDAO dao.MailboxMessagesDAO,
	messageRawBodiesDAO dao.MessageRawBodiesDAO,
	messagesDAO dao.MessagesDAO,
) *SubmissionBackend {
	self := SubmissionBackend{
		accountsDAO:         accountsDAO,
		db:                  db,
		localDelivery:       localDelivery,
		mailboxesDAO:        mailboxesDAO,
		mailboxMessagesDAO:  mailboxMessagesDAO,
		messageRawBodiesDAO: messageRawBodiesDAO,
		messagesDAO:         messagesDAO,
	}
	return &self
}
