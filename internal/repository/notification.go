// Package repository implements persistent data access and processing.
package repository

import (
	"artion-api-graphql/internal/repository/email"
	"artion-api-graphql/internal/types"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// notificationQueueCapacity represents the maximal capacity of notifications queued to be sent.
const notificationQueueCapacity = 100

// NewNotifications provides a channel for new notification requests queue.
func (p *Proxy) NewNotifications() chan types.Notification {
	return p.notifications
}

// QueueNotificationForProcessing puts the given notification into the queue for async processing.
func (p *Proxy) QueueNotificationForProcessing(no *types.Notification) {
	p.notifications <- *no
}

// StoreNotification stores the given notification in persistent storage.
// The function returns true if the notification was unique and didn't exist before.
func (p *Proxy) StoreNotification(no *types.Notification) (bool, error) {
	return p.db.StoreNotification(no)
}

// NotificationTemplates pulls a list of notification templates applicable to the
func (p *Proxy) NotificationTemplates(nt int32, contract *common.Address, tokenID *hexutil.Big) []types.NotificationTemplate {
	return p.db.NotificationTemplates(nt, contract, tokenID)
}

// SendEmailNotificationBySendGrid represents an email notification sender using SendGrid API.
func (p *Proxy) SendEmailNotificationBySendGrid(no *types.Notification, nt *types.NotificationTemplate) error {
	if no == nil || nt == nil {
		return fmt.Errorf("missing notification and/or template")
	}

	// collect recipient user involved
	user, err := p.GetUser(no.Recipient)
	if err != nil {
		return err
	}
	if user == nil {
		return fmt.Errorf("user not found at %s", no.Recipient.String())
	}

	// collect address
	addr, err := p.GetShippingAddress(no.Recipient)
	if err != nil {
		return err
	}

	// send the email
	err = email.SendGridDeliverDynamicTemplate(
		mail.NewEmail(nt.SenderName, nt.SenderID),
		mail.NewEmail("", user.Email),
		nt.TemplateID,
		nt.Subject,
		dynamicTemplateData(no, user, addr, nt.ExtendedParams),
	)
	if err != nil {
		log.Errorf("email notification failed; %s", err.Error())
		return err
	}
	return nil
}

// dynamicTemplateData creates a map of key->value dynamic data points from provided
// source elements.
func dynamicTemplateData(no *types.Notification, usr *types.User, ship *types.ShippingAddress, ext *string) map[string]interface{} {
	var list map[string]interface{}

	// do we have an ext?
	if nil != ext {
		err := json.Unmarshal([]byte(*ext), &list)
		if err != nil {
			log.Errorf("invalid extended data on template; %s", err.Error())
		}
	}

	// add notification details
	if no != nil {
		list["collection"] = no.Contract.String()
		list["contract"] = no.Contract.String()
		list["token_id"] = no.TokenId.String()
	}

	// add user
	if usr != nil {
		list["account"] = usr.Address.String()
		list["address"] = usr.Address.String()
		list["alias"] = usr.Username
		list["email"] = usr.Email
	}

	// add shipping address
	if ship != nil {
		list["ship_name"] = ship.FullName
		list["ship_phone"] = ship.Phone
		list["ship_street"] = ship.Street
		list["ship_apartment"] = ship.Apartment
		list["ship_city"] = ship.City
		list["ship_state"] = ship.State
		list["ship_country"] = ship.Country
		list["ship_zip"] = ship.Zip
	}
	return list
}
