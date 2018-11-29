# golang_design_pattern

<H2>はじめに</H2>
golangを利用したデザインパターンの実装です。

デザインパターンは結城浩さん著作の「Java言語で学ぶデザインパターン入門」を参考にしています。
Javaとgolangは言語仕様が異なるため、Javaでは抽象クラスで表現しているところをgolangでは構造体としたり、
Javaでは継承で表現されているところをgolangではコンポジションに置き換えて設計しています。

<H3>一つ一つ数え上げる Itaratorパターン</H3>

<https://github.com/ShinyaIshikawa/golang_design_pattern/tree/master/iterator>

<H3>皮を被せる Adapterパターン</H3>

<https://github.com/ShinyaIshikawa/golang_design_pattern/tree/master/adapter>

<H3>具体的な処理をサブクラスにまかせる Template Methodパターン</H3>

<https://github.com/ShinyaIshikawa/golang_design_pattern/tree/master/template>
* golangには抽象クラスがないため、コンポジションで再利用を実現しています。
