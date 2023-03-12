from manimlib import *
#manimgl 004Text.py TextExample
class TextExample(Scene):
    def construct(self):
        # 想要正确运行这个场景，你需要确保你的计算机中安装了Consolas字体
        # 关于Text全部用法，请见https://github.com/3b1b/manim/pull/680
        text = Text("Here is a text", font="Consolas", font_size=90)
        difference = Text(
            """
            The most important difference between Text and TexText is that\n
            you can change the font more easily, but can't use the LaTeX grammar

            黄宁真是个天才啊
            """,
            font="Arial", font_size=24,
            # t2c是一个由 文本-颜色 键值对组成的字典
            t2c={"Text": BLUE, "TexText": BLUE, "LaTeX": ORANGE, "黄宁": YELLOW}
        )

        #VGroup 可以将多个 VMobject 放在一起看做一个整体。例子中调用了 arrange() 方法来将其中子物体依次向下排列（DOWN），且间距为 buff
        VGroup(text, difference).arrange(DOWN, buff=1)

        # Write 是显示类似书写效果的动画
        self.play(Write(text))

        # FadeIn 将物体淡入，第二个参数表示淡入的方向
        self.play(Write(difference))
        self.play(FadeIn(difference, UP))

        self.wait(3)

        