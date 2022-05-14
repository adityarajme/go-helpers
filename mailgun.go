package golang_helpers

import (
	"gopkg.in/mailgun/mailgun-go.v1"
)

type Sender struct {
	User     string
	Password string
}

func SendMailGunMail(mailgun_domain string, subject string, content string, mail_handle string, mail_to string, extra_info string, person_name string, file_attachments []string) error {

	mg := mailgun.NewMailgun(mailgun_domain, "key-bfa936211369063e7686bdb16509306f", "pubkey-8e7463e7dad45e74705b129a39026ad9")

	mf := person_name + " <" + mail_handle + "@" + mailgun_domain + ">"
	m_reply_to := person_name + " <" + mail_handle + "@" + mailgun_domain + ">"
	message := mg.NewMessage(mf, subject, "", mail_to)
	message.AddHeader("extra_info", extra_info)
	message.SetReplyTo(m_reply_to)
	message.SetHtml(content)
	if len(file_attachments) > 0 {
		for _, file_attachment := range file_attachments {
			message.AddAttachment("/location_of_files/" + file_attachment)
		}
	}
	_, _, err := mg.Send(message)
	return err
}
