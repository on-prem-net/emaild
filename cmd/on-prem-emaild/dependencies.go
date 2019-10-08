package main

import (
	"github.com/on-prem-net/emaild/cloudservice"
	"github.com/on-prem-net/emaild/dao"
	"github.com/on-prem-net/emaild/imapendpoint"
	"github.com/on-prem-net/emaild/imapendpoint/imapbackend"
	"github.com/on-prem-net/emaild/localdelivery"
	"github.com/on-prem-net/emaild/smtpendpoint"
	"github.com/on-prem-net/emaild/smtpendpoint/smtpbackend"
	"github.com/on-prem-net/emaild/submissionendpoint"
	"github.com/on-prem-net/emaild/submissionendpoint/submissionbackend"
	"github.com/asdine/genji"
	"github.com/asdine/genji/engine"
	imap_backend "github.com/emersion/go-imap/backend"
	"github.com/karlkfi/inject"
)

var (
	graph inject.Graph

	cloudService       *cloudservice.CloudService
	db                 *genji.DB
	genjiEngine        *engine.Engine
	imapBackend        imap_backend.Backend
	imapEndpoint       *imapendpoint.ImapEndpoint
	localDelivery      *localdelivery.LocalDelivery
	smtpBackend        *smtpbackend.SmtpBackend
	smtpEndpoint       *smtpendpoint.SmtpEndpoint
	submissionBackend  *submissionbackend.SubmissionBackend
	submissionEndpoint *submissionendpoint.SubmissionEndpoint

	accountsDAO         dao.AccountsDAO
	mailboxesDAO        dao.MailboxesDAO
	mailboxMessagesDAO  dao.MailboxMessagesDAO
	messageRawBodiesDAO dao.MessageRawBodiesDAO
	messagesDAO         dao.MessagesDAO
	propertiesDAO       dao.PropertiesDAO
)

func DefineDependencies() {
	graph = inject.NewGraph()

	graph.Define(&cloudService, inject.NewAutoProvider(cloudservice.New))
	graph.Define(&db, inject.NewAutoProvider(newDB))
	graph.Define(&genjiEngine, inject.NewAutoProvider(newGenjiEngine))
	graph.Define(&imapBackend, inject.NewAutoProvider(imapbackend.New))
	graph.Define(&imapEndpoint, inject.NewAutoProvider(imapendpoint.New))
	graph.Define(&localDelivery, inject.NewAutoProvider(localdelivery.New))
	graph.Define(&smtpBackend, inject.NewAutoProvider(smtpbackend.New))
	graph.Define(&smtpEndpoint, inject.NewAutoProvider(smtpendpoint.New))
	graph.Define(&submissionBackend, inject.NewAutoProvider(submissionbackend.New))
	graph.Define(&submissionEndpoint, inject.NewAutoProvider(submissionendpoint.New))

	graph.Define(&accountsDAO, inject.NewAutoProvider(dao.NewAccountsDAO))
	graph.Define(&mailboxesDAO, inject.NewAutoProvider(dao.NewMailboxesDAO))
	graph.Define(&mailboxMessagesDAO, inject.NewAutoProvider(dao.NewMailboxMessagesDAO))
	graph.Define(&messageRawBodiesDAO, inject.NewAutoProvider(dao.NewMessageRawBodiesDAO))
	graph.Define(&messagesDAO, inject.NewAutoProvider(dao.NewMessagesDAO))
	graph.Define(&propertiesDAO, inject.NewAutoProvider(dao.NewPropertiesDAO))
}
