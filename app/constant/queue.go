// Copyright 2024 wuzhanfly <wuzhanfly@gmail.com>. All rights reserved.
// Use of this source code is governed by a AGPL v3.0 style
// license that can be found in the LICENSE file.

package constant

const (
	QueueMidjourneyStatusInit       = 0
	QueueMidjourneyStatusProceeding = 1
	QueueMidjourneyStatusEnded      = 2
	QueueMidjourneyStatusError      = 3
)

const (
	QueueMidjourneyEventInsertQueue         = "insertQueue"
	QueueMidjourneyEventInteractionCreate   = "interactionCreate"
	QueueMidjourneyEventInteractionSuccess  = "interactionSuccess"
	QueueMidjourneyEventWriteMessageContent = "writeMessageContent"
	QueueMidjourneyEventEnded               = "ended"
	QueueMidjourneyEventProgress            = "progress"
	QueueMidjourneyEventContinue            = "continue"
	QueueMidjourneyEventVerifyHuman         = "verifyHuman"
	QueueMidjourneyEventError               = "error"
)
