package commands

import (
    "EverythingSuckz/fsb/config"
    "EverythingSuckz/fsb/internal/utils"

    "github.com/celestix/gotgproto/dispatcher"
    "github.com/celestix/gotgproto/dispatcher/handlers"
    "github.com/celestix/gotgproto/ext"
    "github.com/celestix/gotgproto/storage"
)

func (m *command) LoadStart(dispatcher dispatcher.Dispatcher) {
    log := m.log.Named("start")
    defer log.Sugar().Info("Loaded")
    dispatcher.AddHandler(handlers.NewCommand("start", start))
}

func start(ctx *ext.Context, u *ext.Update) error {
    chatId := u.EffectiveChat().GetID()
    peerChatId := ctx.PeerStorage.GetPeerById(chatId)
    if peerChatId.Type != int(storage.TypeUser) {
        return dispatcher.EndGroups
    }
    if len(config.ValueOf.AllowedUsers) != 0 && !utils.Contains(config.ValueOf.AllowedUsers, chatId) {
        ctx.Reply(u, "You are not allowed to use this bot.", nil)
        return dispatcher.EndGroups
    }

    // Force subscribe logic - using a placeholder for ChatMember retrieval
    channelUsername := "@haris_garage" // Replace with your channel username
    isSubscribed := checkSubscription(ctx.Dispatcher(), channelUsername, chatId) // Custom function to handle subscription check
    if !isSubscribed {
        // User is not subscribed
        ctx.Reply(u, "Please join my channel "+channelUsername+" to use this bot.", nil)
        ctx.Reply(u, "Join here: https://t.me/"+channelUsername[1:], nil)
        return dispatcher.EndGroups
    }

    // Normal bot functionality
    ctx.Reply(u, "hi", nil)
    return dispatcher.EndGroups
}

// Placeholder function for subscription check
func checkSubscription(dispatcher dispatcher.Dispatcher, channelUsername string, chatId int64) bool {
    // Implement logic based on how the library provides access to chat members
    // Return true if subscribed, false otherwise
    return true // Replace with actual logic
}
