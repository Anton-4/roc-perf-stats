module Pages.Home_ exposing (view)

import Element exposing (text)
import View exposing (View)


view : View msg
view =
    { title = "Homepage"
    , element = text "Hello, world!"
    }
