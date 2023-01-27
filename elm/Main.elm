module Main exposing (Model, Msg(..), checkbox, init, main, update, view)

import Browser exposing (..)
import Html exposing (..)
import Html.Attributes exposing (..)
import Html.Events exposing (onClick, onInput)



-- MAIN


main =
  Browser.sandbox { init = init, update = update, view = view }



-- MODEL


type alias Model =
  { 
    answer : String 
  , show : Bool
  }


init : Model
init =
  Model "" False



-- UPDATE


type Msg
  = Answer String
  | Show


update : Msg -> Model -> Model
update msg model =
  case msg of
    Answer answer ->
      { model | answer = answer }
    Show ->
      { model | show = not model.show }



-- VIEW


view : Model -> Html Msg
view model =
  div []
  [
  div []
    [ viewTitle model
    , viewInput "text" "Answer" model.answer Answer
    , viewValidation model
    , checkbox Show "show the answer"
    ]
  ]


viewInput : String -> String -> String -> (String -> msg) -> Html msg
viewInput t p v toMsg =
  input [ type_ t, placeholder p, value v, onInput toMsg ] []


viewValidation : Model -> Html msg
viewValidation model =
  if String.toLower(model.answer) == "apple" then
    div [ style "color" "green" ] [ text "You find it !!! " ]
  else
    div [ style "color" "red" ] [ text "wrong word ! " ]

viewTitle : Model -> Html msg
viewTitle model = 
  if model.show == True then
    div [ style "coler" "green" ] [ text "apple" ]
  else
    div [] [ text "Guess the words" ]
    
    
    
    
checkbox : msg -> String -> Html msg
checkbox msg name =
    label
        [ style "padding" "20px" ]
        [ input [ type_ "checkbox", onClick msg ] []
        , text name
        ]
