// Package dasgo provides Type Definitions for the Discord API.
package dasgo

// Gateway Opcodes
// https://discord.com/developers/docs/topics/opcodes-and-status-codes#gateway-gateway-opcodes
const (
	FlagGatewayOpcodeDispatch  = 0
	FlagGatewayOpcodeHeartbeat = 1
	FlagGatewayOpcodeIdentify  = 2
	FlagGatewayOpcodePresence  = 3
	FlagGatewayOpcodeVoice     = 4
	FlagGatewayOpcodeResume    = 6
	FlagGatewayOpcodeReconnect = 7
	FlagGatewayOpcodeRequest   = 8
	FlagGatewayOpcodeInvalid   = 9
	FlagGatewayOpcodeHello     = 10
	FlagGatewayOpcodeAck       = 11
)

// Gateway Close Event Codes
// https://discord.com/developers/docs/topics/opcodes-and-status-codes#gateway-gateway-close-event-codes
const (
	FlagGatewayCloseEventCodeUnknownError         = 4000
	FlagGatewayCloseEventCodeUnknownOpcode        = 4001
	FlagGatewayCloseEventCodeDecodeError          = 4002
	FlagGatewayCloseEventCodeNotAuthenticated     = 4003
	FlagGatewayCloseEventCodeAuthenticationFailed = 4004
	FlagGatewayCloseEventCodeAlreadyAuthenticated = 4005
	FlagGatewayCloseEventCodeInvalidSeq           = 4007
	FlagGatewayCloseEventCodeRateLimited          = 4008
	FlagGatewayCloseEventCodeSessionTimed         = 4009
	FlagGatewayCloseEventCodeInvalidShard         = 4010
	FlagGatewayCloseEventCodeShardingRequired     = 4011
	FlagGatewayCloseEventCodeInvalidAPIVersion    = 4012
	FlagGatewayCloseEventCodeInvalidIntent        = 4013
	FlagGatewayCloseEventCodeDisallowedIntent     = 4014
)

// Voice Opcodes
// https://discord.com/developers/docs/topics/opcodes-and-status-codes#voice-voice-opcodes
const (
	FlagVoiceOpcodeIdentify           = 0
	FlagVoiceOpcodeSelectProtocol     = 1
	FlagVoiceOpcodeReadyServer        = 2
	FlagVoiceOpcodeHeartbeat          = 3
	FlagVoiceOpcodeSessionDescription = 4
	FlagVoiceOpcodeSpeaking           = 5
	FlagVoiceOpcodeHeartbeatACK       = 6
	FlagVoiceOpcodeResume             = 7
	FlagVoiceOpcodeHello              = 8
	FlagVoiceOpcodeResumed            = 9
	FlagVoiceOpcodeClientDisconnect   = 13
)

// Voice Close Event Codes
// https://discord.com/developers/docs/topics/opcodes-and-status-codes#voice-voice-close-event-codes
const (
	FlagVoiceCloseEventCodeUnknownOpcode         = 4001
	FlagVoiceCloseEventCodeFailedDecode          = 4002
	FlagVoiceCloseEventCodeNotAuthenticated      = 4003
	FlagVoiceCloseEventCodeAuthenticationFailed  = 4004
	FlagVoiceCloseEventCodeAlreadyAuthenticated  = 4005
	FlagVoiceCloseEventCodeInvalidSession        = 4006
	FlagVoiceCloseEventCodeSessionTimeout        = 4009
	FlagVoiceCloseEventCodeServerNotFound        = 4011
	FlagVoiceCloseEventCodeUnknownProtocol       = 4012
	FlagVoiceCloseEventCodeDisconnectedChannel   = 4014
	FlagVoiceCloseEventCodeVoiceServerCrash      = 4015
	FlagVoiceCloseEventCodeUnknownEncryptionMode = 4016
)

// HTTP Response Codes
// https://discord.com/developers/docs/topics/opcodes-and-status-codes#http-http-response-codes
const (
	FlagHTTPResponseCodeOK                 = 200
	FlagHTTPResponseCodeCREATED            = 201
	FlagHTTPResponseCodeNOCONTENT          = 204
	FlagHTTPResponseCodeNOTMODIFIED        = 304
	FlagHTTPResponseCodeBADREQUEST         = 400
	FlagHTTPResponseCodeUNAUTHORIZED       = 401
	FlagHTTPResponseCodeFORBIDDEN          = 403
	FlagHTTPResponseCodeNOTFOUND           = 404
	FlagHTTPResponseCodeMETHODNOTALLOWED   = 405
	FlagHTTPResponseCodeTOOMANYREQUESTS    = 429
	FlagHTTPResponseCodeGATEWAYUNAVAILABLE = 502
	FlagHTTPResponseCodeSERVERERROR        = 504 // 5xx (504 Not Guaranteed)
)

// JSON Error Codes
// https://discord.com/developers/docs/topics/opcodes-and-status-codes#json-json-error-codes
var (
	JSONErrorCodes = map[int]string{
		0:      "General error (such as a malformed request body, amongst other things)",
		10001:  "Unknown account",
		10002:  "Unknown application",
		10003:  "Unknown channel",
		10004:  "Unknown guild",
		10005:  "Unknown integration",
		10006:  "Unknown invite",
		10007:  "Unknown member",
		10008:  "Unknown message",
		10009:  "Unknown permission overwrite",
		10010:  "Unknown provider",
		10011:  "Unknown role",
		10012:  "Unknown token",
		10013:  "Unknown user",
		10014:  "Unknown emoji",
		10015:  "Unknown webhook",
		10016:  "Unknown webhook service",
		10020:  "Unknown session",
		10026:  "Unknown ban",
		10027:  "Unknown SKU",
		10028:  "Unknown Store Listing",
		10029:  "Unknown entitlement",
		10030:  "Unknown build",
		10031:  "Unknown lobby",
		10032:  "Unknown branch",
		10033:  "Unknown store directory layout",
		10036:  "Unknown redistributable",
		10038:  "Unknown gift code",
		10049:  "Unknown stream",
		10050:  "Unknown premium server subscribe cooldown",
		10057:  "Unknown guild template",
		10059:  "Unknown discoverable server category",
		10060:  "Unknown sticker",
		10062:  "Unknown interaction",
		10063:  "Unknown application command",
		10065:  "Unknown voice state",
		10066:  "Unknown application command permissions",
		10067:  "Unknown Stage Instance",
		10068:  "Unknown Guild Member Verification Form",
		10069:  "Unknown Guild Welcome Screen",
		10070:  "Unknown Guild Scheduled Event",
		10071:  "Unknown Guild Scheduled Event User",
		20001:  "Bots cannot use this endpoint",
		20002:  "Only bots can use this endpoint",
		20009:  "Explicit content cannot be sent to the desired recipient(s)",
		20012:  "You are not authorized to perform this action on this application",
		20016:  "This action cannot be performed due to slowmode rate limit",
		20018:  "Only the owner of this account can perform this action",
		20022:  "This message cannot be edited due to announcement rate limits",
		20028:  "The channel you are writing has hit the write rate limit",
		20029:  "The write action you are performing on the server has hit the write rate limit",
		20031:  "Your Stage topic, server name, server description, or channel names contain words that are not allowed",
		20035:  "Guild premium subscription level too low",
		30001:  "Maximum number of guilds reached (100)",
		30002:  "Maximum number of friends reached (1000)",
		30003:  "Maximum number of pins reached for the channel (50)",
		30004:  "Maximum number of recipients reached (10)",
		30005:  "Maximum number of guild roles reached (250)",
		30007:  "Maximum number of webhooks reached (10)",
		30008:  "Maximum number of emojis reached",
		30010:  "Maximum number of reactions reached (20)",
		30013:  "Maximum number of guild channels reached (500)",
		30015:  "Maximum number of attachments in a message reached (10)",
		30016:  "Maximum number of invites reached (1000)",
		30018:  "Maximum number of animated emojis reached",
		30019:  "Maximum number of server members reached",
		30030:  "Maximum number of server categories has been reached (5)",
		30031:  "Guild already has a template",
		30033:  "Max number of thread participants has been reached (1000)",
		30035:  "Maximum number of bans for non-guild members have been exceeded",
		30037:  "Maximum number of bans fetches has been reached",
		30038:  "Maximum number of uncompleted guild scheduled events reached (100)",
		30039:  "Maximum number of stickers reached",
		30040:  "Maximum number of prune requests has been reached. Try again later",
		30042:  "Maximum number of guild widget settings updates has been reached. Try again later",
		30046:  "Maximum number of edits to messages older than 1 hour reached. Try again later",
		40001:  "Unauthorized. Provide a valid token and try again",
		40002:  "You need to verify your account in order to perform this action",
		40003:  "You are opening direct messages too fast",
		40004:  "Send messages has been temporarily disabled",
		40005:  "Request entity too large. Try sending something smaller in size",
		40006:  "This feature has been temporarily disabled server-side",
		40007:  "The user is banned from this guild",
		40032:  "Target user is not connected to voice",
		40033:  "This message has already been crossposted",
		40041:  "An application command with that name already exists",
		40060:  "Interaction has already been acknowledged",
		50001:  "Missing access",
		50002:  "Invalid account type",
		50003:  "Cannot execute action on a DM channel",
		50004:  "Guild widget disabled",
		50005:  "Cannot edit a message authored by another user",
		50006:  "Cannot send an empty message",
		50007:  "Cannot send messages to this user",
		50008:  "Cannot send messages in a non-text channel",
		50009:  "Channel verification level is too high for you to gain access",
		50010:  "OAuth2 application does not have a bot",
		50011:  "OAuth2 application limit reached",
		50012:  "Invalid OAuth2 state",
		50013:  "You lack permissions to perform that action",
		50014:  "Invalid authentication token provided",
		50015:  "Note was too long",
		50016:  "Provided too few or too many messages to delete. Must provide at least 2 and fewer than 100 messages to delete",
		50019:  "A message can only be pinned to the channel it was sent in",
		50020:  "Invite code was either invalid or taken",
		50021:  "Cannot execute action on a system message",
		50024:  "Cannot execute action on this channel type",
		50025:  "Invalid OAuth2 access token provided",
		50026:  "Missing required OAuth2 scope",
		50027:  "Invalid webhook token provided",
		50028:  "Invalid role",
		50033:  "Invalid Recipient(s)",
		50034:  "A message provided was too old to bulk delete",
		50035:  "Invalid form body (returned for both application/json and multipart/form-data bodies), or invalid Content-Type provided",
		50036:  "An invite was accepted to a guild the application's bot is not in",
		50041:  "Invalid API version provided",
		50045:  "File uploaded exceeds the maximum size",
		50046:  "Invalid file uploaded",
		50054:  "Cannot self-redeem this gift",
		50055:  "Invalid Guild",
		50068:  "Invalid message type",
		50070:  "Payment source required to redeem gift",
		50074:  "Cannot delete a channel required for Community guilds",
		50081:  "Invalid sticker sent",
		50083:  "Tried to perform an operation on an archived thread, such as editing a message or adding a user to the thread",
		50084:  "Invalid thread notification settings",
		50085:  "before value is earlier than the thread creation date",
		50086:  "Community server channels must be text channels",
		50095:  "This server is not available in your location",
		50097:  "This server needs monetization enabled in order to perform this action",
		50101:  "This server needs more boosts to perform this action",
		50109:  "The request body contains invalid JSON.",
		60003:  "Two factor is required for this operation",
		80004:  "No users with DiscordTag exist",
		90001:  "Reaction was blocked",
		130000: "API resource is currently overloaded. Try again a little later",
		150006: "The Stage is already open",
		160002: "Cannot reply without permission to read message history",
		160004: "A thread has already been created for this message",
		160005: "Thread is locked",
		160006: "Maximum number of active threads reached",
		160007: "Maximum number of active announcement threads reached",
		170001: "Invalid JSON for uploaded Lottie file",
		170002: "Uploaded Lotties cannot contain rasterized images such as PNG or JPEG",
		170003: "Sticker maximum framerate exceeded",
		170004: "Sticker frame count exceeds maximum of 1000 frames",
		170005: "Lottie animation maximum dimensions exceeded",
		170006: "Sticker frame rate is either too small or too large",
		170007: "Sticker animation duration exceeds maximum of 5 seconds",
		180000: "Cannot update a finished event",
		180002: "Failed to create stage needed for stage event",
	}
)

// RPC Error Codes
// https://discord.com/developers/docs/topics/opcodes-and-status-codes#rpc-rpc-error-codes
const (
	FlagRPCErrorCodeUnknownError                    = 1000
	FlagRPCErrorCodeInvalidPayload                  = 4000
	FlagRPCErrorCodeInvalidCommand                  = 4002
	FlagRPCErrorCodeInvalidGuild                    = 4003
	FlagRPCErrorCodeInvalidEvent                    = 4004
	FlagRPCErrorCodeInvalidChannel                  = 4005
	FlagRPCErrorCodeInvalidPermissions              = 4006
	FlagRPCErrorCodeInvalidClientID                 = 4007
	FlagRPCErrorCodeInvalidOrigin                   = 4008
	FlagRPCErrorCodeInvalidToken                    = 4009
	FlagRPCErrorCodeInvalidUser                     = 4010
	FlagRPCErrorCodeOAuth2Error                     = 5000
	FlagRPCErrorCodeSelectChannelTimedOut           = 5001
	FlagRPCErrorCodeGET_GUILDTimedOut               = 5002
	FlagRPCErrorCodeSelectVoiceForceRequired        = 5003
	FlagRPCErrorCodeCaptureShortcutAlreadyListening = 5004
)

// RPC Close Event Codes
// https://discord.com/developers/docs/topics/opcodes-and-status-codes#rpc-rpc-close-event-codes
const (
	FlagRPCCloseEventCodeInvalidClientID = 4000
	FlagRPCCloseEventCodeInvalidOrigin   = 4001
	FlagRPCCloseEventCodeRateLimited     = 4002
	FlagRPCCloseEventCodeTokenRevoked    = 4003
	FlagRPCCloseEventCodeInvalidVersion  = 4004
	FlagRPCCloseEventCodeInvalidEncoding = 4005
)
