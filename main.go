package main

import (
	"flag"
	"fmt"
	"os"
)

const binary = "slackcli"
const version = "1.0.6"

func main() {
	cli := NewCLI(binary)

	cli.AutoAuthenticate()

	cli.Register(cli.CallAPITest, "api.test", []string{"error"}, "Checks API calling code")
	cli.Register(cli.CallAppsConnectionsOpen, "apps.connections.open", []string{}, "Generate a temporary Socket Mode WebSocket URL that your app can connect to in order to receive events and interactive payloads over")
	cli.Register(cli.CallAppsEventAuthorizationsList, "apps.event.authorizations.list", []string{"event_context", "cursor", "limit"}, "Get a list of authorizations for the given event context. Each authorization represents an app installation that the event is visible to")
	cli.Register(cli.CallAppsList, "apps.list", []string{}, "Lists associated applications")
	cli.Register(cli.CallAppsManifestCreate, "apps.manifest.create", []string{"manifest"}, "Create an app from an app manifest")
	cli.Register(cli.CallAppsManifestDelete, "apps.manifest.delete", []string{"app_id"}, "Permanently deletes an app created through app manifests")
	cli.Register(cli.CallAppsManifestExport, "apps.manifest.export", []string{"app_id"}, "Export an app manifest from an existing app")
	cli.Register(cli.CallAppsManifestUpdate, "apps.manifest.update", []string{"app_id", "manifest"}, "Update an app from an app manifest")
	cli.Register(cli.CallAppsManifestValidate, "apps.manifest.validate", []string{"manifest", "app_id"}, "Validate an app manifest")
	cli.Register(cli.CallAuthRevoke, "auth.revoke", []string{"test"}, "Revokes a token")
	cli.Register(cli.CallAuthTeamsList, "auth.teams.list", []string{"cursor", "include_icon", "limit"}, "List the workspaces a token can access")
	cli.Register(cli.CallAuthTest, "auth.test", []string{}, "Checks authentication and identity")
	cli.Register(cli.CallBotsInfo, "bots.info", []string{"bot"}, "Gets information about a bot user")
	cli.Register(cli.CallChatDelete, "chat.delete", []string{"channel", "time"}, "Deletes a message")
	cli.Register(cli.CallChatDeleteAttachment, "chat.deleteAttachment", []string{"channel", "time", "attachment"}, "Deletes a message attachment")
	cli.Register(cli.CallChatMeMessage, "chat.meMessage", []string{"channel", "text"}, "Share a me message into a channel")
	cli.Register(cli.CallChatPostAttachment, "chat.postAttachment", []string{"channel", "json"}, "Sends an attachment to a channel")
	cli.Register(cli.CallChatPostMessage, "chat.postMessage", []string{"channel", "text"}, "Sends a message to a channel")
	cli.Register(cli.CallChatRobotMessage, "chat.robotMessage", []string{"channel", "text"}, "Sends a message to a channel as a robot")
	cli.Register(cli.CallChatUpdate, "chat.update", []string{"channel", "time", "text"}, "Updates a message")
	cli.Register(cli.CallConversationsAcceptSharedInvite, "conversations.acceptSharedInvite", []string{"channel_name", "channel_id", "free_trial_accepted", "invite_id", "is_private", "team_id"}, "Accepts an invitation to a Slack Connect channel")
	cli.Register(cli.CallConversationsApproveSharedInvite, "conversations.approveSharedInvite", []string{"invite_id", "target_team"}, "Approves an invitation to a Slack Connect channel")
	cli.Register(cli.CallConversationsArchive, "conversations.archive", []string{"room"}, "Archives a conversation")
	cli.Register(cli.CallConversationsClose, "conversations.close", []string{"room"}, "Closes a direct message or multi-person direct message")
	cli.Register(cli.CallConversationsCreate, "conversations.create", []string{"name", "is_private", "team_id"}, "Initiates a public or private channel-based conversation")
	cli.Register(cli.CallConversationsDeclineSharedInvite, "conversations.declineSharedInvite", []string{"invite_id", "target_team"}, "Declines a Slack Connect channel invite")
	cli.Register(cli.CallConversationsGenericInfo, "conversations.genericInfo", []string{"channels"}, "Retrieve information about various channels")
	cli.Register(cli.CallConversationsHistory, "conversations.history", []string{"room", "time"}, "Fetches a conversation's history of messages and events")
	cli.Register(cli.CallConversationsInfo, "conversations.info", []string{"room"}, "Retrieve information about a conversation")
	cli.Register(cli.CallConversationsInvite, "conversations.invite", []string{"room", "user"}, "Invites users to a channel")
	cli.Register(cli.CallConversationsInviteShared, "conversations.inviteShared", []string{"channel", "emails", "external_limited", "user_ids"}, "Sends an invitation to a Slack Connect channel")
	cli.Register(cli.CallConversationsJoin, "conversations.join", []string{"room"}, "Joins an existing conversation")
	cli.Register(cli.CallConversationsKick, "conversations.kick", []string{"room", "user"}, "Removes a user from a conversation")
	cli.Register(cli.CallConversationsLeave, "conversations.leave", []string{"room"}, "Leaves a conversation")
	cli.Register(cli.CallConversationsList, "conversations.list", []string{}, "Lists all channels in a Slack team")
	cli.Register(cli.CallConversationsListConnectInvites, "conversations.listConnectInvites", []string{"count", "cursor", "team_id"}, "Lists shared channel invites that have been generated or received but have not been approved by all parties")
	cli.Register(cli.CallConversationsMark, "conversations.mark", []string{"room", "time"}, "Sets the read cursor in a channel")
	cli.Register(cli.CallConversationsMembers, "conversations.members", []string{"channel", "cursor", "limit"}, "Retrieve members of a conversation")
	cli.Register(cli.CallConversationsOpen, "conversations.open", []string{"channel", "prevent_creation", "return_im", "users"}, "Opens or resumes a direct message or multi-person direct message")
	cli.Register(cli.CallConversationsRename, "conversations.rename", []string{"room", "name"}, "Renames a conversation")
	cli.Register(cli.CallConversationsReplies, "conversations.replies", []string{"channel", "ts", "cursor", "inclusive", "latest", "limit", "oldest"}, "Retrieve a thread of messages posted to a conversation")
	cli.Register(cli.CallConversationsSetPurpose, "conversations.setPurpose", []string{"room", "purpose"}, "Sets the purpose for a conversation")
	cli.Register(cli.CallConversationsSetTopic, "conversations.setTopic", []string{"room", "topic"}, "Sets the topic for a conversation")
	cli.Register(cli.CallConversationsUnarchive, "conversations.unarchive", []string{"room"}, "Reverses conversation archival")
	cli.Register(cli.CallDndEndDnd, "dnd.endDnd", []string{}, "Ends the current user's \"Do Not Disturb\" session immediately")
	cli.Register(cli.CallDndEndSnooze, "dnd.endSnooze", []string{}, "Ends the current user's snooze mode immediately")
	cli.Register(cli.CallDndInfo, "dnd.info", []string{"user"}, "Retrieves a user's current \"Do Not Disturb\" status")
	cli.Register(cli.CallDndSetSnooze, "dnd.setSnooze", []string{}, "Ends the current user's snooze mode immediately")
	cli.Register(cli.CallDndTeamInfo, "dnd.teamInfo", []string{"users"}, "Retrieves the \"Do Not Disturb\" status for users on a team")
	cli.Register(cli.CallEmojiList, "emoji.list", []string{}, "Lists custom emoji for a team")
	cli.Register(cli.CallEventlogHistory, "eventlog.history", []string{"time"}, "Lists all the events since the specified time")
	cli.Register(cli.CallFilesCommentsAdd, "files.comments.add", []string{"file", "text"}, "Add a comment to an existing file")
	cli.Register(cli.CallFilesCommentsDelete, "files.comments.delete", []string{"file", "fcid"}, "Deletes an existing comment on a file")
	cli.Register(cli.CallFilesCommentsEdit, "files.comments.edit", []string{"file", "fcid", "text"}, "Edit an existing file comment")
	cli.Register(cli.CallFilesDelete, "files.delete", []string{"file"}, "Deletes a file and associated comments")
	cli.Register(cli.CallFilesInfo, "files.info", []string{"file", "count", "page"}, "Gets information about a team file")
	cli.Register(cli.CallFilesList, "files.list", []string{"count", "page"}, "Lists and filters team files")
	cli.Register(cli.CallFilesListAfterTime, "files.listAfterTime", []string{"time", "count", "page"}, "Lists and filters team files after this timestamp (inclusive)")
	cli.Register(cli.CallFilesListBeforeTime, "files.listBeforeTime", []string{"time", "count", "page"}, "Lists and filters team files before this timestamp (inclusive)")
	cli.Register(cli.CallFilesListByChannel, "files.listByChannel", []string{"channel", "count", "page"}, "Lists and filters team files in a specific channel")
	cli.Register(cli.CallFilesListByType, "files.listByType", []string{"type", "count", "page"}, "Lists and filters team files by type: all, posts, snippets, images, gdocs, zips, pdfs")
	cli.Register(cli.CallFilesListByUser, "files.listByUser", []string{"user", "count", "page"}, "Lists and filters team files created by a single user")
	cli.Register(cli.CallFilesRevokePublicURL, "files.revokePublicURL", []string{"file"}, "Revokes public/external sharing access for a file")
	cli.Register(cli.CallFilesSharedPublicURL, "files.sharedPublicURL", []string{"file"}, "Enables a file for public/external sharing")
	cli.Register(cli.CallFilesUpload, "files.upload", []string{"channel", "filename"}, "Uploads or creates a file from local data")
	cli.Register(cli.CallHelpIssuesList, "help.issues.list", []string{}, "List issues reported by the current user")
	cli.Register(cli.CallMigrationExchange, "migration.exchange", []string{"users", "order"}, "For Enterprise Grid workspaces, map local user IDs to global user IDs")
	cli.Register(cli.CallPinsAdd, "pins.add", []string{"channel", "item_id"}, "Pins an item to a channel")
	cli.Register(cli.CallPinsList, "pins.list", []string{"channel"}, "Lists items pinned to a channel")
	cli.Register(cli.CallPinsRemove, "pins.remove", []string{"channel", "item_id"}, "Un-pins an item from a channel")
	cli.Register(cli.CallReactionsAdd, "reactions.add", []string{"channel", "time", "name"}, "Adds a reaction to an item")
	cli.Register(cli.CallReactionsGet, "reactions.get", []string{"channel", "time"}, "Gets reactions for an item")
	cli.Register(cli.CallReactionsList, "reactions.list", []string{"user"}, "Lists reactions made by a user")
	cli.Register(cli.CallReactionsRemove, "reactions.remove", []string{"channel", "time", "name"}, "Removes a reaction from an item")
	cli.Register(cli.CallRtmEvents, "rtm.events", []string{}, "Prints the API events in real time")
	cli.Register(cli.CallSignupCheckEmail, "signup.checkEmail", []string{"email"}, "Checks if an email address is valid")
	cli.Register(cli.CallSignupConfirmEmail, "signup.confirmEmail", []string{"email"}, "Confirm an email address for signup")
	cli.Register(cli.CallSearchAll, "search.all", []string{"query", "count", "page"}, "Searches for messages and files matching a query")
	cli.Register(cli.CallSearchFiles, "search.files", []string{"query", "count", "page"}, "Searches for files matching a query")
	cli.Register(cli.CallSearchMessages, "search.messages", []string{"query", "count", "page"}, "Searches for messages matching a query")
	cli.Register(cli.CallSearchUsers, "search.users", []string{"user", "count"}, "Search users by name or email address")
	cli.Register(cli.CallStarsAdd, "stars.add", []string{"channel", "item_id"}, "Adds a star to an item")
	cli.Register(cli.CallStarsList, "stars.list", []string{"count", "page"}, "Lists stars for a user")
	cli.Register(cli.CallStarsRemove, "stars.remove", []string{"channel", "item_id"}, "Removes a star from an item")
	cli.Register(cli.CallTeamAccessLogs, "team.accessLogs", []string{"before", "count", "page"}, "Gets the access logs for the current team")
	cli.Register(cli.CallTeamBillableInfo, "team.billableInfo", []string{"team_id", "user"}, "Gets billable users information for the current team")
	cli.Register(cli.CallTeamBillingInfo, "team.billing.info", []string{}, "Reads a workspace's billing plan information")
	cli.Register(cli.CallTeamInfo, "team.info", []string{"team"}, "Gets information about the current team")
	cli.Register(cli.CallTeamIntegrationLogs, "team.integrationLogs", []string{"app_id", "change_type", "count", "page", "service_id", "team_id", "user"}, "Gets the integration logs for the current team")
	cli.Register(cli.CallTeamPreferencesList, "team.preferences.list", []string{}, "Retrieve a list of a workspace's team preferences")
	cli.Register(cli.CallTeamProfileGet, "team.profile.get", []string{}, "Retrieve a team's profile")
	cli.Register(cli.CallUsersCounts, "users.counts", []string{}, "Count number of users in the team")
	cli.Register(cli.CallUsersDeletePhoto, "users.deletePhoto", []string{}, "Delete the user avatar")
	cli.Register(cli.CallUsersGetPresence, "users.getPresence", []string{"user"}, "Gets user presence information")
	cli.Register(cli.CallUsersID, "users.id", []string{"user", "limit"}, "Gets user identifier from username")
	cli.Register(cli.CallUsersIdentity, "users.identity", []string{}, "Get a user's identity")
	cli.Register(cli.CallUsersInfo, "users.info", []string{"user"}, "Gets information about a user")
	cli.Register(cli.CallUsersList, "users.list", []string{"limit", "cursor"}, "Lists all users in a Slack team")
	cli.Register(cli.CallUsersLookupByEmail, "users.lookupByEmail", []string{"email"}, "Find a user with an email address")
	cli.Register(cli.CallUsersPrefsGet, "users.prefs.get", []string{}, "Get user account preferences")
	cli.Register(cli.CallUsersPrefsSet, "users.prefs.set", []string{"name", "value"}, "Set user account preferences")
	cli.Register(cli.CallUsersPreparePhoto, "users.preparePhoto", []string{"image"}, "Upload a picture to use as the avatar")
	cli.Register(cli.CallUsersProfileGet, "users.profile.get", []string{"user"}, "Retrieves a user's profile information")
	cli.Register(cli.CallUsersProfileSet, "users.profile.set", []string{"name", "value"}, "Set the profile information for a user")
	cli.Register(cli.CallUsersSetActive, "users.setActive", []string{}, "Marks a user as active")
	cli.Register(cli.CallUsersSetAvatar, "users.setAvatar", []string{"image"}, "Upload a picture and set it as the avatar")
	cli.Register(cli.CallUsersSetEmail, "users.setEmail", []string{"email"}, "Changes the email address without confirmation")
	cli.Register(cli.CallUsersSetPhoto, "users.setPhoto", []string{"image_id"}, "Define which picture will be the avatar")
	cli.Register(cli.CallUsersSetPresence, "users.setPresence", []string{"presence"}, "Manually sets user presence")
	cli.Register(cli.CallUsersSetStatus, "users.setStatus", []string{"emoji", "text"}, "Set the status message and emoji")
	cli.Register(cli.CallUsersSetUsername, "users.setUsername", []string{"username"}, "Changes the username without admin privileges")
	cli.Register(cli.CallWorkflowsStepCompleted, "workflows.stepCompleted", []string{"workflow_step_execute_id"}, "Indicate that an app's step in a workflow completed execution")
	cli.Register(cli.CallWorkflowsStepFailed, "workflows.stepFailed", []string{"workflow_step_execute_id", "error"}, "Indicate that an app's step in a workflow failed to execute")
	cli.Register(cli.CallWorkflowsUpdateStep, "workflows.updateStep", []string{"workflow_step_edit_id", "step_image_url", "step_name"}, "Update the configuration for a workflow step")
	cli.Register(cli.CallVersion, "version", []string{}, "Displays the program version number")
	cli.Register(cli.CallHelp, "help", []string{}, "Displays usage and program options")

	flag.Usage = func() {
		fmt.Println("Slack CLI v" + version)
		fmt.Println("  http://cixtor.com/")
		fmt.Println("  https://api.slack.com/")
		fmt.Println("  https://github.com/cixtor/slackapi")
		fmt.Println("  https://github.com/cixtor/slackcli")
		fmt.Println()
		fmt.Println("Description:")
		fmt.Println("  Low level Slack API client with custom commands. Slack, the 'messaging app for")
		fmt.Println("  teams' offers an API that has been used to build multiple projects around it,")
		fmt.Println("  from bots to independent clients as well as integrations with other external")
		fmt.Println("  services. This project aims to offer a low level experience for advanced users")
		fmt.Println("  that want to either drop the web client or interact with the API for testing")
		fmt.Println("  purpose.")
		fmt.Println()
		fmt.Println("Usage:")

		cli.PrintCommands()
	}

	flag.Parse()

	os.Exit(cli.Execute(flag.Arg(0)))
}
