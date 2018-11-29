<H3>具体的な処理をサブクラスにまかせる Template Methodパターン</H3>

オリジナルは抽象クラスに処理の大枠を定義して、具体的な処理をサブクラスに任せていますが、
本ソースではdisplay_contorollerで処理の大枠を定義して、具体的な処理はdisplayにまかせています。

* display_contoroller  
処理の大枠を定めています。 display_contorollerはdisplayを持ちます。

* display(interface)
表示に関する振る舞いを定義しています。

* char_display
displayをimplemetします。

* string_display
displayをimplemetします。
