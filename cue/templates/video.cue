package templates

import (
        "strings"
)

#Video: {
  id: string
  url: *"https://youtu.be/" | string
  title: string & strings.MaxRunes(70)
  description: *"N/A" | string
  likes: *0 | int & >= 0
}

idp: #Video & {
    id: "Rg98GoEHBd4"
    url: "https://youtu.be/" + id
    title: "How To Create A Complete Internal Developer Platform (IDP)?"
    description: "It's time to build an internal developer platform (IDO) with Crossplane, Argo CD, SchemaHero, External Secrets Operator (ESO), GitHub Actions, Port, and a few others."
    likes: 12532
}

top2023: #Video & {
    id: "-zKYT2I_WhE"
    url: "https://youtu.be/" + id
    title: "The Best DevOps Tools, Platforms, And Services In 2023?"
    likes: 21341
}

top2022: #Video & {
    id: "7wIr6e6b7Ps"
    url: "https://youtu.be/" + id
    title: "The Best DevOps Tools, Frameworks, And Platforms In 2022"
    likes: 12341
}
