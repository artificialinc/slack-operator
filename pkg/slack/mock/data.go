package mock

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/slack-go/slack"
)

var ConversationName = "bat-channel"
var PublicConversationID = "C0EAQDV4Z"
var PrivateConversationID = "Y7HGFWC6Q"
var BotID = "U023BECGF"
var Description = "My channel Description"

var templateChannelJSON = `
{
	"id": "%s",
	"name": "%s",
	"is_channel": true,
	"is_group": false,
	"is_im": false,
	"created": %d,
	"creator": "%s",
	"is_archived": false,
	"is_general": false,
	"unlinked": 0,
	"name_normalized": "%s",
	"is_shared": false,
	"is_ext_shared": false,
	"is_org_shared": false,
	"pending_shared": [],
	"is_pending_ext_shared": false,
	"is_member": true,
	"is_private": %s,
	"is_mpim": false,
	"last_read": "0000000000.000000",
	"latest": null,
	"unread_count": 0,
	"unread_count_display": 0,
	"topic": {
		"value": "%s",
		"creator": "%s",
		"last_set": %d
	},
	"purpose": {
		"value": "%s",
		"creator": "%s",
		"last_set": %d
	},
	"num_members": %d,
	"previous_names": [],
	"priority": 0
}`

var templateConversationJSON = fmt.Sprintf(`
	{
		"ok": true,
		"channel": %s
	}`, templateChannelJSON)

var publicConversationJSON = fmt.Sprintf(templateConversationJSON, PublicConversationID, ConversationName,
	nowAsJSONTime(), BotID, ConversationName, "false", "", "", 0, "", "", 0, 0)

var privateConversationJSON = fmt.Sprintf(templateConversationJSON, PrivateConversationID, ConversationName,
	nowAsJSONTime(), BotID, ConversationName, "true", "", "", 0, "", "", 0, 0)

var inviteConversationJSON = fmt.Sprintf(templateConversationJSON, PublicConversationID, ConversationName,
	nowAsJSONTime(), BotID, ConversationName, "false", "", "", 0, "", "", 0, 1)

func getConversationNameResponse(name string) string {
	return fmt.Sprintf(templateConversationJSON, PublicConversationID, name,
		nowAsJSONTime(), BotID, name, "false", "", "", 0, "", "", 0, 0)
}

func getConversationTopicResponse(topic string) string {
	return fmt.Sprintf(templateConversationJSON, PublicConversationID, ConversationName,
		nowAsJSONTime(), BotID, ConversationName, "false", topic, BotID,
		nowAsJSONTime(), "I didn't set this purpose on purpose!", BotID, nowAsJSONTime(), 0)
}

func getConversationPurposeResponse(purpose string) string {
	return fmt.Sprintf(templateConversationJSON, PublicConversationID, ConversationName,
		nowAsJSONTime(), BotID, ConversationName, "false", "random topic", BotID,
		nowAsJSONTime(), purpose, BotID, nowAsJSONTime(), 0)
}

func getChannelPurposeResponse(purpose string) string {
	return fmt.Sprintf(templateChannelJSON, PublicConversationID, ConversationName,
		nowAsJSONTime(), BotID, ConversationName, "false", "random topic", BotID,
		nowAsJSONTime(), purpose, BotID, nowAsJSONTime(), 0)
}

const userEmail = "spengler@ghostbusters.example.com"

var templateUserJSON = `
{
    "ok": true,
    "user": {
        "id": "W012A3CDE",
        "team_id": "T012AB3C4",
        "name": "spengler",
        "deleted": false,
        "color": "9f69e7",
        "real_name": "Egon Spengler",
        "tz": "America/Los_Angeles",
        "tz_label": "Pacific Daylight Time",
        "tz_offset": -25200,
        "profile": {
            "avatar_hash": "ge3b51ca72de",
            "status_text": "Print is dead",
            "status_emoji": ":books:",
            "real_name": "Egon Spengler",
            "display_name": "spengler",
            "real_name_normalized": "Egon Spengler",
            "display_name_normalized": "spengler",
            "email": "%s",
            "image_24": "https://.../avatar/e3b51ca72dee4ef87916ae2b9240df50.jpg",
            "image_32": "https://.../avatar/e3b51ca72dee4ef87916ae2b9240df50.jpg",
            "image_48": "https://.../avatar/e3b51ca72dee4ef87916ae2b9240df50.jpg",
            "image_72": "https://.../avatar/e3b51ca72dee4ef87916ae2b9240df50.jpg",
            "image_192": "https://.../avatar/e3b51ca72dee4ef87916ae2b9240df50.jpg",
            "image_512": "https://.../avatar/e3b51ca72dee4ef87916ae2b9240df50.jpg",
            "team": "T012AB3C4"
        },
        "is_admin": true,
        "is_owner": false,
        "is_primary_owner": false,
        "is_restricted": false,
        "is_ultra_restricted": false,
        "is_bot": false,
        "updated": 1502138686,
        "is_app_user": false,
        "has_2fa": false
    }
}`

var userJSON = fmt.Sprintf(templateUserJSON, userEmail)

func nowAsJSONTime() slack.JSONTime {
	return slack.JSONTime(time.Now().Unix())
}

func asChannelObject(j string) *slack.Channel {
	channel := &slack.Channel{}
	err := json.Unmarshal([]byte(j), &channel)

	if err != nil {
		panic(err)
	}
	return channel
}