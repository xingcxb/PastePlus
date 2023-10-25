//go:build darwin

package application

import "github.com/wailsapp/wails/v3/pkg/events"

var commonApplicationEventMap = map[events.ApplicationEventType]events.ApplicationEventType{
	events.Mac.ApplicationDidFinishLaunching: events.Common.ApplicationStarted,
	events.Mac.ApplicationDidChangeTheme:     events.Common.ThemeChanged,
}

func (m *macosApp) setupCommonEvents() {
	for sourceEvent, targetEvent := range commonApplicationEventMap {
		sourceEvent := sourceEvent
		targetEvent := targetEvent
		m.parent.On(sourceEvent, func(event *Event) {
			event.Id = uint(targetEvent)
			applicationEvents <- event
		})
	}
}
