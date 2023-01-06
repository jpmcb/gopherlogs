package emojis

const (
	// The following are a set of emoji codes that can be used
	// with the logger.Event and logger.Eventf logging methods in the log package.
	//
	// For the purposes of this gopherlogs.emojis package, only a few are included
	// as constants for the sake of simplicity and to reduce the runtime package size.
	//
	// If you want to use your own emojis, find their string value here:
	// https://www.unicode.org/emoji/charts/full-emoji-list.html
	//
	// Generally, a chosen emoji should take up only 2 terminal columns.
	WrenchEmoji     = "\U0001F527"
	FolderEmoji     = "\U0001F4C1"
	PictureEmoji    = "\U0001F3A8"
	PackageEmoji    = "\U0001F4E6"
	RocketEmoji     = "\U0001F680"
	EnvelopeEmoji   = "\U0001F4E7"
	GlobeEmoji      = "\U0001F310"
	GreenCheckEmoji = "\U00002705"
	ControllerEmoji = "\U0001F3AE"
	TestTubeEmoji   = "\U0001F9EA"
	MagnetEmoji     = "\U0001F9F2"
)
