rm -rf doc
mkdir -p doc/model
mkdir -p doc/directive/display/template
mkdir -p doc/directive/audio_player
mkdir -p doc/directive/video_player
mkdir -p doc/card

cd bot

godoc2md github.com/dueros/bot-sdk-go/bot Bot > ../doc/bot.md
godoc2md github.com/dueros/bot-sdk-go/bot Application > ../doc/application.md
godoc2md github.com/dueros/bot-sdk-go/bot/model Dialog > ../doc/model/dialog.md
godoc2md github.com/dueros/bot-sdk-go/bot/model Intent> ../doc/model/intent.md

godoc2md github.com/dueros/bot-sdk-go/bot/model Request> ../doc/model/request.md
godoc2md github.com/dueros/bot-sdk-go/bot/model LaunchRequest>> ../doc/model/request.md
godoc2md github.com/dueros/bot-sdk-go/bot/model IntentRequest>> ../doc/model/request.md
godoc2md github.com/dueros/bot-sdk-go/bot/model SessionEndedRequest>> ../doc/model/request.md
godoc2md github.com/dueros/bot-sdk-go/bot/model EventRequest>> ../doc/model/request.md

godoc2md github.com/dueros/bot-sdk-go/bot/model Response> ../doc/model/response.md
godoc2md github.com/dueros/bot-sdk-go/bot/model Session> ../doc/model/session.md
godoc2md github.com/dueros/bot-sdk-go/bot/model SSMLTextBuilder> ../doc/model/ssml-text-builder.md

godoc2md github.com/dueros/bot-sdk-go/bot/card > ../doc/card/card.md

godoc2md github.com/dueros/bot-sdk-go/bot/directive/display/template > ../doc/directive/display/template.md
godoc2md github.com/dueros/bot-sdk-go/bot/directive/display RenderTemplate> ../doc/directive/display/render-template.md
godoc2md github.com/dueros/bot-sdk-go/bot/directive/audio_player > ../doc/directive/audio-player.md
godoc2md github.com/dueros/bot-sdk-go/bot/directive/video_player > ../doc/directive/video-player.md

cd -

