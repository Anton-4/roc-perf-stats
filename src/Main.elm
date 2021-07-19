module Main exposing (..)

import Browser
import Html exposing (Html, div, h1, text)



---- MODEL ----


type alias Model =
    {}


init : ( Model, Cmd Msg )
init =
    ( {}, Cmd.none )



---- UPDATE ----


type Msg
    = NoOp


update : Msg -> Model -> ( Model, Cmd Msg )
update _ model =
    ( model, Cmd.none )



---- VIEW ----


view : Model -> Html Msg
view _ =
    div []
        [ h1 [] [ text "Elm App is working!" ]
        ]



---- PROGRAM ---


main : Program () Model Msg
main =
    Browser.document
        { view = \model -> { title = "Roc Perf Stats", body = [ view model ] }
        , init = \_ -> init
        , update = update
        , subscriptions = always Sub.none
        }
