## Go中的抽象工厂设计模式

抽象工厂设计模式是一种创新的设计模式，可让您创建一系列相关的对象。它是对工厂模式的抽象。最好用一个例

子来解释。假设我们有两个工厂:

- nike
- adidas

想象一下，你需要买一套运动服，它有一只鞋，但很短。 大多数时候，你最好是买一套类似工厂的全套运动装备

比如耐克或阿迪达斯。 这就是抽象工厂进入画面的地方，你想要的具体产品是鞋子和短裤，这些产品将由耐克和

阿迪达斯的抽象工厂创造。

耐克和阿迪达斯这两家工厂都实现了iSportsFactory接口我们有两个产品界面。

- IShoe-此接口由nikeShoe和adidasShoe实现

- IShort-此接口由nikeShort和adidasShort实现。

## 实现

- `iSportsFactory.go` 定义抽象工厂,根据传入的参数返回不同的工厂实例(nike或者adidas)

```go
package main

import "fmt"

type iSportsFactory interface {
    makeShoe() iShoe
    makeShort() iShort
}

func getSportsFactory(brand string) (iSportsFactory, error) {
    if brand == "adidas" {
        return &adidas{}, nil
    }
    if brand == "nike" {
        return &nike{}, nil
    }
    return nil, fmt.Errorf("Wrong brand type passed")
}
```
- `adidas` 实现 `iSportsFactory` 接口

```go
type adidas struct {}

func (a *adidas) makeShoe() iShoe {
    return &adidasShoe{
        shoe: shoe{
            logo: "adidas",
            size: 14,
        },
    }
}

func (a *adidas) makeShort() iShort {
    return &adidasShort{
        short: short{
            logo: "adidas",
            size: 14,
        },
    }
}
```

- `nike` 实现 `iSportsFactory` 接口

```go
type nike struct {}

func (n *nike) makeShoe() iShoe {
    return &nikeShoe{
        shoe: shoe{
            logo: "nike",
            size: 14,
        },
    }
}

func (n *nike) makeShort() iShort {
    return &nikeShort{
        short: short{
            logo: "nike",
            size: 14,
        },
    }
}
```

- iShow

```go
type iShoe interface {
    setLogo(logo string)
    setSize(size int)
    getLogo() string
    getSize() int
}

type shoe struct {
    logo string
    size int
}

func (s *shoe) setLogo(logo string) {
    s.logo = logo
}

func (s *shoe) getLogo() string {
    return s.logo
}

func (s *shoe) setSize(size int) {
    s.size = size
}

func (s *shoe) getSize() int {
    return s.size
}
```
