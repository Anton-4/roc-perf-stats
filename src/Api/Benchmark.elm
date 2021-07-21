module Api.Benchmark exposing (Benchmark)

import Time


type alias Benchmark =
    { id : String
    , name : String
    , lowBoundConfInterval : Float
    , lowBoundUnit : String
    , pointEstimate : Float
    , pointEstimateUnit : String
    , upBoundConfInterval : Float
    , upBoundUnit : String
    , nrOfMeasurements : Int
    , nrOfOutliers : Int
    , startTime : Time.Posix
    , commit : String
    , branch : String
    , createdAt : Time.Posix
    }
