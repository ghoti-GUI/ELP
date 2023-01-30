-- Make a GET request to load a book called "Public Opinion"
--
-- Read how it works:
--   https://guide.elm-lang.org/effects/http.html
--
module Main exposing (Model, Msg(..), init, main, update, view)

import Browser exposing (..)
import Html exposing (..)
import Html.Attributes exposing (..)
import Html.Events exposing (onClick, onInput)
import Http



-- MAIN


main =
  Browser.element
    { init = init
    , update = update
    , subscriptions = subscriptions
    , view = view
    }



-- MODEL


type alias Model = 
  {
    text:String
  , answer:String
  , show:Bool
  , signal:Signal
  }
  
type Signal
  = Failure
  | Loading
  | Success


init : () -> (Model, Cmd Msg)
init _ =
  ( Model "" "" False Loading
  , Http.get
      { url = "https://elm-lang.org/assets/public-opinion.txt"
      , expect = Http.expectString GotText
      }
  )



-- UPDATE


type Msg
  = GotText (Result Http.Error String)
  | Answer String
  | Show


update : Msg -> Model -> (Model, Cmd Msg)
update msg model =
  case msg of
    GotText result ->
      case result of
        Ok fullText ->
          ({ model | text = fullText, signal = Success}, Cmd.none)

        Err _ ->
          ({model | signal = Failure}, Cmd.none)
    Answer usranswer ->
      ({model | answer=usranswer}, Cmd.none)
    Show ->
      ({model | show = not model.show}, Cmd.none)



-- SUBSCRIPTIONS


subscriptions : Model -> Sub Msg
subscriptions model =
  Sub.none



-- VIEW


view : Model -> Html Msg
view model =
  case model.signal of
    Failure ->
      text "I was unable to load your book."

    Loading ->
      text "Loading..."

    Success ->
      div[]
      [viewTitle model
      , viewInput "text" "Answer" model.answer Answer
      , viewValidation model
      , checkbox Show "show the answer"
      , pre [] [ text model.text ]
      ]




viewTitle : Model -> Html msg
viewTitle model = 
  if model.show == True then
    div [ style "coler" "green" ] [ text "apple" ]
  else
    div [] [ text "Guess the words" ]
    
   
viewInput : String -> String -> String -> (String -> msg) -> Html msg
viewInput t p v toMsg =
  input [ type_ t, placeholder p, value v, onInput toMsg ] []
  
  
viewValidation : Model -> Html msg
viewValidation model =
  if String.toLower(model.answer) == "apple" then
    div [ style "color" "green" ] [ text "You find it !!! " ]
  else
    div [ style "color" "red" ] [ text "wrong word ! " ]
    

checkbox : msg -> String -> Html msg
checkbox msg name =
    label
        [ style "padding" "20px" ]
        [ input [ type_ "checkbox", onClick msg ] []
        , text name
        ]