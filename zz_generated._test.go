// This file has been automatically generated. Don't edit it.

package goobs_test

import (
	"net/http"
	"os"
	"testing"

	goobs "github.com/andreykaipov/goobs"
	config "github.com/andreykaipov/goobs/api/requests/config"
	filters "github.com/andreykaipov/goobs/api/requests/filters"
	general "github.com/andreykaipov/goobs/api/requests/general"
	inputs "github.com/andreykaipov/goobs/api/requests/inputs"
	mediainputs "github.com/andreykaipov/goobs/api/requests/media_inputs"
	outputs "github.com/andreykaipov/goobs/api/requests/outputs"
	rconfig "github.com/andreykaipov/goobs/api/requests/rconfig"
	record "github.com/andreykaipov/goobs/api/requests/record"
	sceneitems "github.com/andreykaipov/goobs/api/requests/scene_items"
	scenes "github.com/andreykaipov/goobs/api/requests/scenes"
	sources "github.com/andreykaipov/goobs/api/requests/sources"
	stream "github.com/andreykaipov/goobs/api/requests/stream"
	transitions "github.com/andreykaipov/goobs/api/requests/transitions"
	ui "github.com/andreykaipov/goobs/api/requests/ui"
	typedefs "github.com/andreykaipov/goobs/api/typedefs"
	assert "github.com/stretchr/testify/assert"
)

func Test_config(t *testing.T) {
	client, err := goobs.New(
		"localhost:"+os.Getenv("OBS_PORT"),
		goobs.WithPassword("goodpassword"),
		goobs.WithRequestHeader(http.Header{"User-Agent": []string{"goobs-e2e/0.0.0"}}),
	)
	assert.NoError(t, err)
	t.Cleanup(func() {
		client.Disconnect()
	})

	_, err = client.Config.CreateProfile(&config.CreateProfileParams{ProfileName: "test"})
	assert.NoError(t, err)
	_, err = client.Config.CreateSceneCollection(&config.CreateSceneCollectionParams{SceneCollectionName: "test"})
	assert.Error(t, err)
	_, err = client.Config.GetPersistentData(&config.GetPersistentDataParams{
		Realm:    "OBS_WEBSOCKET_DATA_REALM_GLOBAL",
		SlotName: "test",
	})
	assert.NoError(t, err)
	_, err = client.Config.GetProfileList(&config.GetProfileListParams{})
	assert.NoError(t, err)
	_, err = client.Config.GetProfileParameter(&config.GetProfileParameterParams{
		ParameterCategory: "test",
		ParameterName:     "test",
	})
	assert.NoError(t, err)
	_, err = client.Config.GetSceneCollectionList(&config.GetSceneCollectionListParams{})
	assert.NoError(t, err)
	_, err = client.Config.GetStreamServiceSettings(&config.GetStreamServiceSettingsParams{})
	assert.NoError(t, err)
	_, err = client.Config.GetVideoSettings(&config.GetVideoSettingsParams{})
	assert.NoError(t, err)
	_, err = client.Config.SetCurrentProfile(&config.SetCurrentProfileParams{ProfileName: "test"})
	assert.NoError(t, err)
	_, err = client.Config.SetCurrentSceneCollection(
		&config.SetCurrentSceneCollectionParams{SceneCollectionName: "test"},
	)
	assert.NoError(t, err)
	_, err = client.Config.SetPersistentData(&config.SetPersistentDataParams{
		Realm:     "OBS_WEBSOCKET_DATA_REALM_GLOBAL",
		SlotName:  "test",
		SlotValue: "",
	})
	assert.NoError(t, err)
	_, err = client.Config.SetProfileParameter(&config.SetProfileParameterParams{
		ParameterCategory: "test",
		ParameterName:     "test",
		ParameterValue:    "test",
	})
	assert.NoError(t, err)
	_, err = client.Config.SetStreamServiceSettings(&config.SetStreamServiceSettingsParams{
		StreamServiceSettings: &typedefs.StreamServiceSettings{},
		StreamServiceType:     "rtmp_custom",
	})
	assert.NoError(t, err)
	_, err = client.Config.SetVideoSettings(&config.SetVideoSettingsParams{
		BaseHeight:     10.0,
		BaseWidth:      10.0,
		FpsDenominator: 10.0,
		FpsNumerator:   10.0,
		OutputHeight:   10.0,
		OutputWidth:    10.0,
	})
	assert.NoError(t, err)
	_, err = client.Config.RemoveProfile(&config.RemoveProfileParams{ProfileName: "test"})
	assert.NoError(t, err)
}

func Test_filters(t *testing.T) {
	client, err := goobs.New(
		"localhost:"+os.Getenv("OBS_PORT"),
		goobs.WithPassword("goodpassword"),
		goobs.WithRequestHeader(http.Header{"User-Agent": []string{"goobs-e2e/0.0.0"}}),
	)
	assert.NoError(t, err)
	t.Cleanup(func() {
		client.Disconnect()
	})

	_, err = client.Filters.CreateSourceFilter(&filters.CreateSourceFilterParams{
		FilterKind:     "scroll_filter",
		FilterName:     "test",
		FilterSettings: map[string]interface{}{"test": "test"},
		SourceName:     "test",
	})
	assert.NoError(t, err)
	_, err = client.Filters.GetSourceFilterDefaultSettings(
		&filters.GetSourceFilterDefaultSettingsParams{FilterKind: "scroll_filter"},
	)
	assert.NoError(t, err)
	_, err = client.Filters.GetSourceFilterList(&filters.GetSourceFilterListParams{SourceName: "test"})
	assert.NoError(t, err)
	_, err = client.Filters.GetSourceFilter(&filters.GetSourceFilterParams{
		FilterName: "test",
		SourceName: "test",
	})
	assert.NoError(t, err)
	_, err = client.Filters.SetSourceFilterEnabled(&filters.SetSourceFilterEnabledParams{
		FilterEnabled: &[]bool{true}[0],
		FilterName:    "test",
		SourceName:    "test",
	})
	assert.NoError(t, err)
	_, err = client.Filters.SetSourceFilterIndex(&filters.SetSourceFilterIndexParams{
		FilterIndex: 1.0,
		FilterName:  "test",
		SourceName:  "test",
	})
	assert.NoError(t, err)
	_, err = client.Filters.SetSourceFilterName(&filters.SetSourceFilterNameParams{
		FilterName:    "test",
		NewFilterName: "test",
		SourceName:    "test",
	})
	assert.Error(t, err)
	_, err = client.Filters.SetSourceFilterSettings(&filters.SetSourceFilterSettingsParams{
		FilterName:     "test",
		FilterSettings: map[string]interface{}{"test": "test"},
		Overlay:        &[]bool{true}[0],
		SourceName:     "test",
	})
	assert.NoError(t, err)
	_, err = client.Filters.RemoveSourceFilter(&filters.RemoveSourceFilterParams{
		FilterName: "test",
		SourceName: "test",
	})
	assert.NoError(t, err)
}

func Test_general(t *testing.T) {
	client, err := goobs.New(
		"localhost:"+os.Getenv("OBS_PORT"),
		goobs.WithPassword("goodpassword"),
		goobs.WithRequestHeader(http.Header{"User-Agent": []string{"goobs-e2e/0.0.0"}}),
	)
	assert.NoError(t, err)
	t.Cleanup(func() {
		client.Disconnect()
	})

	_, err = client.General.BroadcastCustomEvent(
		&general.BroadcastCustomEventParams{EventData: map[string]bool{"test": true}},
	)
	assert.NoError(t, err)
	_, err = client.General.CallVendorRequest(&general.CallVendorRequestParams{
		RequestData: "",
		RequestType: "test",
		VendorName:  "test",
	})
	assert.Error(t, err)
	_, err = client.General.GetHotkeyList(&general.GetHotkeyListParams{})
	assert.NoError(t, err)
	_, err = client.General.GetStats(&general.GetStatsParams{})
	assert.NoError(t, err)
	_, err = client.General.GetVersion(&general.GetVersionParams{})
	assert.NoError(t, err)
	_, err = client.General.Sleep(&general.SleepParams{
		SleepFrames: 1.0,
		SleepMillis: 1.0,
	})
	assert.Error(t, err)
	_, err = client.General.TriggerHotkeyByKeySequence(&general.TriggerHotkeyByKeySequenceParams{
		KeyId:        "OBS_KEY_SHIFT",
		KeyModifiers: &typedefs.KeyModifiers{},
	})
	assert.NoError(t, err)
	_, err = client.General.TriggerHotkeyByName(
		&general.TriggerHotkeyByNameParams{HotkeyName: "OBSBasic.ShowContextBar"},
	)
	assert.NoError(t, err)
}

func Test_inputs(t *testing.T) {
	client, err := goobs.New(
		"localhost:"+os.Getenv("OBS_PORT"),
		goobs.WithPassword("goodpassword"),
		goobs.WithRequestHeader(http.Header{"User-Agent": []string{"goobs-e2e/0.0.0"}}),
	)
	assert.NoError(t, err)
	t.Cleanup(func() {
		client.Disconnect()
	})

	_, err = client.Inputs.CreateInput(&inputs.CreateInputParams{
		InputKind:        "ffmpeg_source",
		InputName:        "test2",
		InputSettings:    map[string]interface{}{"test": "test"},
		SceneItemEnabled: &[]bool{true}[0],
		SceneName:        "Scene",
	})
	assert.NoError(t, err)
	_, err = client.Inputs.GetInputAudioBalance(&inputs.GetInputAudioBalanceParams{InputName: "test2"})
	assert.NoError(t, err)
	_, err = client.Inputs.GetInputAudioMonitorType(&inputs.GetInputAudioMonitorTypeParams{InputName: "test2"})
	assert.NoError(t, err)
	_, err = client.Inputs.GetInputAudioSyncOffset(&inputs.GetInputAudioSyncOffsetParams{InputName: "test2"})
	assert.NoError(t, err)
	_, err = client.Inputs.GetInputAudioTracks(&inputs.GetInputAudioTracksParams{InputName: "test2"})
	assert.NoError(t, err)
	_, err = client.Inputs.GetInputDefaultSettings(&inputs.GetInputDefaultSettingsParams{InputKind: "ffmpeg_source"})
	assert.NoError(t, err)
	_, err = client.Inputs.GetInputKindList(&inputs.GetInputKindListParams{Unversioned: &[]bool{true}[0]})
	assert.NoError(t, err)
	_, err = client.Inputs.GetInputList(&inputs.GetInputListParams{InputKind: "ffmpeg_source"})
	assert.NoError(t, err)
	_, err = client.Inputs.GetInputMute(&inputs.GetInputMuteParams{InputName: "test2"})
	assert.NoError(t, err)
	_, err = client.Inputs.GetInputPropertiesListPropertyItems(&inputs.GetInputPropertiesListPropertyItemsParams{
		InputName:    "test2",
		PropertyName: "test2",
	})
	assert.Error(t, err)
	_, err = client.Inputs.GetInputSettings(&inputs.GetInputSettingsParams{InputName: "test2"})
	assert.NoError(t, err)
	_, err = client.Inputs.GetInputVolume(&inputs.GetInputVolumeParams{InputName: "test2"})
	assert.NoError(t, err)
	_, err = client.Inputs.GetSpecialInputs(&inputs.GetSpecialInputsParams{})
	assert.NoError(t, err)
	_, err = client.Inputs.PressInputPropertiesButton(&inputs.PressInputPropertiesButtonParams{
		InputName:    "test2",
		PropertyName: "test2",
	})
	assert.Error(t, err)
	_, err = client.Inputs.SetInputAudioBalance(&inputs.SetInputAudioBalanceParams{
		InputAudioBalance: 1.0,
		InputName:         "test2",
	})
	assert.NoError(t, err)
	_, err = client.Inputs.SetInputAudioMonitorType(&inputs.SetInputAudioMonitorTypeParams{
		InputName:   "test2",
		MonitorType: "test2",
	})
	assert.Error(t, err)
	_, err = client.Inputs.SetInputAudioSyncOffset(&inputs.SetInputAudioSyncOffsetParams{
		InputAudioSyncOffset: 1.0,
		InputName:            "test2",
	})
	assert.NoError(t, err)
	_, err = client.Inputs.SetInputAudioTracks(&inputs.SetInputAudioTracksParams{
		InputAudioTracks: &typedefs.InputAudioTracks{"test": true},
		InputName:        "test2",
	})
	assert.NoError(t, err)
	_, err = client.Inputs.SetInputMute(&inputs.SetInputMuteParams{
		InputMuted: &[]bool{true}[0],
		InputName:  "test2",
	})
	assert.NoError(t, err)
	_, err = client.Inputs.SetInputName(&inputs.SetInputNameParams{
		InputName:    "test2",
		NewInputName: "test2",
	})
	assert.Error(t, err)
	_, err = client.Inputs.SetInputSettings(&inputs.SetInputSettingsParams{
		InputName:     "test2",
		InputSettings: map[string]interface{}{"test": "test"},
		Overlay:       &[]bool{true}[0],
	})
	assert.NoError(t, err)
	_, err = client.Inputs.SetInputVolume(&inputs.SetInputVolumeParams{
		InputName:      "test2",
		InputVolumeDb:  1.0,
		InputVolumeMul: 1.0,
	})
	assert.Error(t, err)
	_, err = client.Inputs.ToggleInputMute(&inputs.ToggleInputMuteParams{InputName: "test2"})
	assert.NoError(t, err)
	_, err = client.Inputs.RemoveInput(&inputs.RemoveInputParams{InputName: "test2"})
	assert.NoError(t, err)
}

func Test_media_inputs(t *testing.T) {
	client, err := goobs.New(
		"localhost:"+os.Getenv("OBS_PORT"),
		goobs.WithPassword("goodpassword"),
		goobs.WithRequestHeader(http.Header{"User-Agent": []string{"goobs-e2e/0.0.0"}}),
	)
	assert.NoError(t, err)
	t.Cleanup(func() {
		client.Disconnect()
	})

	_, err = client.MediaInputs.GetMediaInputStatus(&mediainputs.GetMediaInputStatusParams{InputName: "test"})
	assert.NoError(t, err)
	_, err = client.MediaInputs.OffsetMediaInputCursor(&mediainputs.OffsetMediaInputCursorParams{
		InputName:         "test",
		MediaCursorOffset: 1.0,
	})
	assert.NoError(t, err)
	_, err = client.MediaInputs.SetMediaInputCursor(&mediainputs.SetMediaInputCursorParams{
		InputName:   "test",
		MediaCursor: 1.0,
	})
	assert.NoError(t, err)
	_, err = client.MediaInputs.TriggerMediaInputAction(&mediainputs.TriggerMediaInputActionParams{
		InputName:   "test",
		MediaAction: "OBS_WEBSOCKET_MEDIA_INPUT_ACTION_PAUSE",
	})
	assert.NoError(t, err)
}

func Test_outputs(t *testing.T) {
	client, err := goobs.New(
		"localhost:"+os.Getenv("OBS_PORT"),
		goobs.WithPassword("goodpassword"),
		goobs.WithRequestHeader(http.Header{"User-Agent": []string{"goobs-e2e/0.0.0"}}),
	)
	assert.NoError(t, err)
	t.Cleanup(func() {
		client.Disconnect()
	})

	_, err = client.Outputs.GetLastReplayBufferReplay(&outputs.GetLastReplayBufferReplayParams{})
	assert.Error(t, err)
	_, err = client.Outputs.GetReplayBufferStatus(&outputs.GetReplayBufferStatusParams{})
	assert.Error(t, err)
	_, err = client.Outputs.GetVirtualCamStatus(&outputs.GetVirtualCamStatusParams{})
	assert.Error(t, err)
	_, err = client.Outputs.SaveReplayBuffer(&outputs.SaveReplayBufferParams{})
	assert.Error(t, err)
	_, err = client.Outputs.StartReplayBuffer(&outputs.StartReplayBufferParams{})
	assert.Error(t, err)
	_, err = client.Outputs.StartVirtualCam(&outputs.StartVirtualCamParams{})
	assert.Error(t, err)
	_, err = client.Outputs.StopReplayBuffer(&outputs.StopReplayBufferParams{})
	assert.Error(t, err)
	_, err = client.Outputs.StopVirtualCam(&outputs.StopVirtualCamParams{})
	assert.Error(t, err)
	_, err = client.Outputs.ToggleReplayBuffer(&outputs.ToggleReplayBufferParams{})
	assert.Error(t, err)
	_, err = client.Outputs.ToggleVirtualCam(&outputs.ToggleVirtualCamParams{})
	assert.Error(t, err)
}

func Test_rconfig(t *testing.T) {
	client, err := goobs.New(
		"localhost:"+os.Getenv("OBS_PORT"),
		goobs.WithPassword("goodpassword"),
		goobs.WithRequestHeader(http.Header{"User-Agent": []string{"goobs-e2e/0.0.0"}}),
	)
	assert.NoError(t, err)
	t.Cleanup(func() {
		client.Disconnect()
	})

	_, err = client.Rconfig.GetRecordDirectory(&rconfig.GetRecordDirectoryParams{})
	assert.NoError(t, err)
}

func Test_record(t *testing.T) {
	client, err := goobs.New(
		"localhost:"+os.Getenv("OBS_PORT"),
		goobs.WithPassword("goodpassword"),
		goobs.WithRequestHeader(http.Header{"User-Agent": []string{"goobs-e2e/0.0.0"}}),
	)
	assert.NoError(t, err)
	t.Cleanup(func() {
		client.Disconnect()
	})

	_, err = client.Record.GetRecordStatus(&record.GetRecordStatusParams{})
	assert.NoError(t, err)
	_, err = client.Record.PauseRecord(&record.PauseRecordParams{})
	assert.NoError(t, err)
	_, err = client.Record.ResumeRecord(&record.ResumeRecordParams{})
	assert.Error(t, err)
	_, err = client.Record.StartRecord(&record.StartRecordParams{})
	assert.NoError(t, err)
	_, err = client.Record.StopRecord(&record.StopRecordParams{})
	assert.Error(t, err)
	_, err = client.Record.ToggleRecord(&record.ToggleRecordParams{})
	assert.NoError(t, err)
	_, err = client.Record.ToggleRecordPause(&record.ToggleRecordPauseParams{})
	assert.NoError(t, err)
}

func Test_scene_items(t *testing.T) {
	client, err := goobs.New(
		"localhost:"+os.Getenv("OBS_PORT"),
		goobs.WithPassword("goodpassword"),
		goobs.WithRequestHeader(http.Header{"User-Agent": []string{"goobs-e2e/0.0.0"}}),
	)
	assert.NoError(t, err)
	t.Cleanup(func() {
		client.Disconnect()
	})

	_, err = client.SceneItems.CreateSceneItem(&sceneitems.CreateSceneItemParams{
		SceneItemEnabled: &[]bool{true}[0],
		SceneName:        "Scene",
		SourceName:       "test",
	})
	assert.NoError(t, err)
	_, err = client.SceneItems.DuplicateSceneItem(&sceneitems.DuplicateSceneItemParams{
		DestinationSceneName: "test",
		SceneItemId:          1.0,
		SceneName:            "Scene",
	})
	assert.Error(t, err)
	_, err = client.SceneItems.GetGroupItemList(&sceneitems.GetGroupItemListParams{SceneName: "Scene"})
	assert.Error(t, err)
	_, err = client.SceneItems.GetSceneItemBlendMode(&sceneitems.GetSceneItemBlendModeParams{
		SceneItemId: 1.0,
		SceneName:   "Scene",
	})
	assert.NoError(t, err)
	_, err = client.SceneItems.GetSceneItemEnabled(&sceneitems.GetSceneItemEnabledParams{
		SceneItemId: 1.0,
		SceneName:   "Scene",
	})
	assert.NoError(t, err)
	_, err = client.SceneItems.GetSceneItemId(&sceneitems.GetSceneItemIdParams{
		SceneName:    "Scene",
		SearchOffset: 1.0,
		SourceName:   "test",
	})
	assert.NoError(t, err)
	_, err = client.SceneItems.GetSceneItemIndex(&sceneitems.GetSceneItemIndexParams{
		SceneItemId: 1.0,
		SceneName:   "Scene",
	})
	assert.NoError(t, err)
	_, err = client.SceneItems.GetSceneItemList(&sceneitems.GetSceneItemListParams{SceneName: "Scene"})
	assert.NoError(t, err)
	_, err = client.SceneItems.GetSceneItemLocked(&sceneitems.GetSceneItemLockedParams{
		SceneItemId: 1.0,
		SceneName:   "Scene",
	})
	assert.NoError(t, err)
	_, err = client.SceneItems.GetSceneItemTransform(&sceneitems.GetSceneItemTransformParams{
		SceneItemId: 1.0,
		SceneName:   "Scene",
	})
	assert.NoError(t, err)
	_, err = client.SceneItems.SetSceneItemBlendMode(&sceneitems.SetSceneItemBlendModeParams{
		SceneItemBlendMode: "OBS_BLEND_NORMAL",
		SceneItemId:        1.0,
		SceneName:          "Scene",
	})
	assert.NoError(t, err)
	_, err = client.SceneItems.SetSceneItemEnabled(&sceneitems.SetSceneItemEnabledParams{
		SceneItemEnabled: &[]bool{true}[0],
		SceneItemId:      1.0,
		SceneName:        "Scene",
	})
	assert.NoError(t, err)
	_, err = client.SceneItems.SetSceneItemIndex(&sceneitems.SetSceneItemIndexParams{
		SceneItemId:    1.0,
		SceneItemIndex: 1.0,
		SceneName:      "Scene",
	})
	assert.NoError(t, err)
	_, err = client.SceneItems.SetSceneItemLocked(&sceneitems.SetSceneItemLockedParams{
		SceneItemId:     1.0,
		SceneItemLocked: &[]bool{true}[0],
		SceneName:       "Scene",
	})
	assert.NoError(t, err)
	_, err = client.SceneItems.SetSceneItemTransform(&sceneitems.SetSceneItemTransformParams{
		SceneItemId: 1.0,
		SceneItemTransform: &typedefs.SceneItemTransform{
			BoundsHeight: 1.0,
			BoundsType:   "OBS_BOUNDS_NONE",
			BoundsWidth:  1.0,
		},
		SceneName: "Scene",
	})
	assert.NoError(t, err)
	_, err = client.SceneItems.RemoveSceneItem(&sceneitems.RemoveSceneItemParams{
		SceneItemId: 1.0,
		SceneName:   "nonexistent",
	})
	assert.Error(t, err)
}

func Test_scenes(t *testing.T) {
	client, err := goobs.New(
		"localhost:"+os.Getenv("OBS_PORT"),
		goobs.WithPassword("goodpassword"),
		goobs.WithRequestHeader(http.Header{"User-Agent": []string{"goobs-e2e/0.0.0"}}),
	)
	assert.NoError(t, err)
	t.Cleanup(func() {
		client.Disconnect()
	})

	_, err = client.Scenes.CreateScene(&scenes.CreateSceneParams{SceneName: "Scene"})
	assert.Error(t, err)
	_, err = client.Scenes.GetCurrentPreviewScene(&scenes.GetCurrentPreviewSceneParams{})
	assert.NoError(t, err)
	_, err = client.Scenes.GetCurrentProgramScene(&scenes.GetCurrentProgramSceneParams{})
	assert.NoError(t, err)
	_, err = client.Scenes.GetGroupList(&scenes.GetGroupListParams{})
	assert.NoError(t, err)
	_, err = client.Scenes.GetSceneList(&scenes.GetSceneListParams{})
	assert.NoError(t, err)
	_, err = client.Scenes.GetSceneSceneTransitionOverride(
		&scenes.GetSceneSceneTransitionOverrideParams{SceneName: "Scene"},
	)
	assert.NoError(t, err)
	_, err = client.Scenes.SetCurrentPreviewScene(&scenes.SetCurrentPreviewSceneParams{SceneName: "Scene"})
	assert.NoError(t, err)
	_, err = client.Scenes.SetCurrentProgramScene(&scenes.SetCurrentProgramSceneParams{SceneName: "Scene"})
	assert.NoError(t, err)
	_, err = client.Scenes.SetSceneName(&scenes.SetSceneNameParams{
		NewSceneName: "Scene",
		SceneName:    "Scene",
	})
	assert.Error(t, err)
	_, err = client.Scenes.SetSceneSceneTransitionOverride(&scenes.SetSceneSceneTransitionOverrideParams{
		SceneName:          "Scene",
		TransitionDuration: 50.0,
		TransitionName:     "Cut",
	})
	assert.NoError(t, err)
	_, err = client.Scenes.RemoveScene(&scenes.RemoveSceneParams{SceneName: "nonexistent"})
	assert.Error(t, err)
}

func Test_sources(t *testing.T) {
	client, err := goobs.New(
		"localhost:"+os.Getenv("OBS_PORT"),
		goobs.WithPassword("goodpassword"),
		goobs.WithRequestHeader(http.Header{"User-Agent": []string{"goobs-e2e/0.0.0"}}),
	)
	assert.NoError(t, err)
	t.Cleanup(func() {
		client.Disconnect()
	})

	_, err = client.Sources.GetSourceActive(&sources.GetSourceActiveParams{SourceName: "test"})
	assert.NoError(t, err)
	_, err = client.Sources.GetSourceScreenshot(&sources.GetSourceScreenshotParams{
		ImageCompressionQuality: 8.0,
		ImageFormat:             "png",
		ImageHeight:             8.0,
		ImageWidth:              8.0,
		SourceName:              "test",
	})
	assert.NoError(t, err)
	_, err = client.Sources.SaveSourceScreenshot(&sources.SaveSourceScreenshotParams{
		ImageCompressionQuality: 8.0,
		ImageFilePath:           "test",
		ImageFormat:             "png",
		ImageHeight:             8.0,
		ImageWidth:              8.0,
		SourceName:              "test",
	})
	assert.NoError(t, err)
}

func Test_stream(t *testing.T) {
	client, err := goobs.New(
		"localhost:"+os.Getenv("OBS_PORT"),
		goobs.WithPassword("goodpassword"),
		goobs.WithRequestHeader(http.Header{"User-Agent": []string{"goobs-e2e/0.0.0"}}),
	)
	assert.NoError(t, err)
	t.Cleanup(func() {
		client.Disconnect()
	})

	_, err = client.Stream.GetStreamStatus(&stream.GetStreamStatusParams{})
	assert.NoError(t, err)
	_, err = client.Stream.SendStreamCaption(&stream.SendStreamCaptionParams{CaptionText: "test"})
	assert.Error(t, err)
	_, err = client.Stream.StartStream(&stream.StartStreamParams{})
	assert.NoError(t, err)
	_, err = client.Stream.StopStream(&stream.StopStreamParams{})
	assert.Error(t, err)
	_, err = client.Stream.ToggleStream(&stream.ToggleStreamParams{})
	assert.NoError(t, err)
}

func Test_transitions(t *testing.T) {
	client, err := goobs.New(
		"localhost:"+os.Getenv("OBS_PORT"),
		goobs.WithPassword("goodpassword"),
		goobs.WithRequestHeader(http.Header{"User-Agent": []string{"goobs-e2e/0.0.0"}}),
	)
	assert.NoError(t, err)
	t.Cleanup(func() {
		client.Disconnect()
	})

	_, err = client.Transitions.GetCurrentSceneTransitionCursor(&transitions.GetCurrentSceneTransitionCursorParams{})
	assert.NoError(t, err)
	_, err = client.Transitions.GetCurrentSceneTransition(&transitions.GetCurrentSceneTransitionParams{})
	assert.NoError(t, err)
	_, err = client.Transitions.GetSceneTransitionList(&transitions.GetSceneTransitionListParams{})
	assert.NoError(t, err)
	_, err = client.Transitions.GetTransitionKindList(&transitions.GetTransitionKindListParams{})
	assert.NoError(t, err)
	_, err = client.Transitions.SetCurrentSceneTransitionDuration(
		&transitions.SetCurrentSceneTransitionDurationParams{TransitionDuration: 50.0},
	)
	assert.NoError(t, err)
	_, err = client.Transitions.SetCurrentSceneTransition(
		&transitions.SetCurrentSceneTransitionParams{TransitionName: "Cut"},
	)
	assert.NoError(t, err)
	_, err = client.Transitions.SetCurrentSceneTransitionSettings(&transitions.SetCurrentSceneTransitionSettingsParams{
		Overlay:            &[]bool{true}[0],
		TransitionSettings: map[string]interface{}{"test": "test"},
	})
	assert.Error(t, err)
	_, err = client.Transitions.SetTBarPosition(&transitions.SetTBarPositionParams{
		Position: 1.0,
		Release:  &[]bool{true}[0],
	})
	assert.NoError(t, err)
	_, err = client.Transitions.TriggerStudioModeTransition(&transitions.TriggerStudioModeTransitionParams{})
	assert.NoError(t, err)
}

func Test_ui(t *testing.T) {
	client, err := goobs.New(
		"localhost:"+os.Getenv("OBS_PORT"),
		goobs.WithPassword("goodpassword"),
		goobs.WithRequestHeader(http.Header{"User-Agent": []string{"goobs-e2e/0.0.0"}}),
	)
	assert.NoError(t, err)
	t.Cleanup(func() {
		client.Disconnect()
	})

	_, err = client.Ui.GetMonitorList(&ui.GetMonitorListParams{})
	assert.NoError(t, err)
	_, err = client.Ui.GetStudioModeEnabled(&ui.GetStudioModeEnabledParams{})
	assert.NoError(t, err)
	_, err = client.Ui.OpenInputFiltersDialog(&ui.OpenInputFiltersDialogParams{InputName: "test"})
	assert.NoError(t, err)
	_, err = client.Ui.OpenInputInteractDialog(&ui.OpenInputInteractDialogParams{InputName: "test"})
	assert.Error(t, err)
	_, err = client.Ui.OpenInputPropertiesDialog(&ui.OpenInputPropertiesDialogParams{InputName: "test"})
	assert.NoError(t, err)
	_, err = client.Ui.SetStudioModeEnabled(&ui.SetStudioModeEnabledParams{StudioModeEnabled: &[]bool{true}[0]})
	assert.NoError(t, err)
}
