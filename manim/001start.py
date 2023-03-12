from manimlib import *

class SquareToCircle(Scene):
    def construct(self):
        circle = Circle()
        circle.set_fill(BLUE, opacity=0.5)
        circle.set_stroke(BLUE_E, width=4)
        square = Square()

        self.play(ShowCreation(square))
        self.wait(3)
        self.play(ReplacementTransform(square, circle)) #ReplacementTransform(A, B) 表示把A转换为B的图案并替代B
        self.wait(2)


        self.embed() #支持交互是新版本的新特性