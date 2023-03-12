from manimlib import *
class TextExample(Scene):
    def construct(self):
        fonts = Text(
            "And you can also set the font according to different words",
            font="Arial",
            t2f={"font": "Consolas", "words": "Consolas"},
            t2c={"font": BLUE, "words": GREEN}
        )
        fonts.set_width(FRAME_WIDTH - 1)

        slant = Text(
            "And the same as slant and weight",
            font="Consolas",
            t2s={"slant": ITALIC},
            t2w={"weight": BOLD},
            t2c={"slant": ORANGE, "weight": RED}
        )

        VGroup(fonts, slant).arrange(DOWN, buff=0.8)
        # you can attemp to combine the write and fadeout into one play function
        # some thing strange will happen...
        self.play(Write(fonts), Write(slant))
        self.play(FadeOut(fonts, shift=DOWN), FadeOut(slant, shift=UP))
        self.wait(2)
        self.wait(1)