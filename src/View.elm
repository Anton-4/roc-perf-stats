module View exposing (View, map, none, placeholder, toBrowserDocument)

import Browser
import Element exposing (Element, layout, text)


type alias View msg =
    { title : String
    , element : Element msg
    }


placeholder : String -> View msg
placeholder pageName =
    { title = pageName
    , element = text pageName
    }


none : View msg
none =
    { title = ""
    , element = Element.none
    }


map : (a -> b) -> View a -> View b
map fn view =
    { title = view.title
    , element = Element.map fn view.element
    }


toBrowserDocument : View msg -> Browser.Document msg
toBrowserDocument view =
    { title = view.title
    , body =
        [ layout [] view.element
        ]
    }
