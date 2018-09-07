# release-point

A sample for the release operation of touches using the [ebiten](https://github.com/hajimehoshi/ebiten) on a smartphone.

## How to work

1. Save the ID by `inpututil.JustPressedTouchIDs()`
2. Track every frame position using the ID until released.
3. Confirm the release by `inpututil.IsTouchJustReleased(t.ID)`.

## How to build

After connecting the debugging Android smartphone to the PC, execute the following.

```
gomobile install github.com/kemokemo/ebiten-sketchbook/release-point
```
