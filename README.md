# part2d

part2d is a simple particle physics engine on a 2D plane.

## Table of contents

1. [Project structure](#project-structure)

## Project structure

```
.
├── .editorconfig
├── .gitignore
├── go.mod
├── LICENSE
├── pkg
│   └── part2d
│       └── part2d.go
└── README.md
```

## Info

part2d uses integers instead of floats. This means that you get to choose your precision level by providing all values represented in the smallest unit that is significant for your application. For example, gravity (-9.81 m/s<sup>2</sup>) would need to be rounded to either -9 m/s<sup>2</sup> or -10 m/s<sup>2</sup>. If your application needs more precision, then it might be better to represent gravity as -981 cm/s<sup>2</sup> or -9810 mm/s<sup>2</sup>. Just make sure that all of your other values are scaled appropriately as well.
