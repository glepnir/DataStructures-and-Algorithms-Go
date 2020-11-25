# 创建者模式

实例创建可以很简单，例如提供左括号和右括号{}，并为实例保留零值，也可以复杂到需要进行一些API调用，检查状态并为其字段创建对象的对象。 您还可以拥有一个由许多对象组成的对象，这在Go中确实很常用，因为它不支持继承。

同时，您可能使用相同的技术来创建多种类型的对象。 例如，您将使用与制造公共汽车几乎相同的技术来制造汽车，除了它们的大小和座位数不同之外，为什么我们不重复使用制造过程呢？ 这就是构建器模式的得力之处。

## 目的

一个创建者模式试图做到以下几点:

* 抽象复杂的创建，以便对象创建与对象用户分离

* 通过填充其字段并创建嵌入的对象来逐步创建对象

* 在许多对象之间重用对象创建算法

## 示例-车辆制造

通常将Builder设计模式描述为导演，一些Builders和他们生产的产品之间的关系。以汽车为例，我们将创建一个汽车制造商，每种车辆的创建车辆（产品）的过程（被广泛称为算法）都大致相同-选择车辆类型，组装结构，放置车轮和放置座椅。 如果考虑一

下，您可以使用此描述来制造汽车和摩托车（两个制造商），因此我们将重复使用该描述来制造制造中的汽车。 在我们的示例中，导演由ManufacturingDirector类型表示.

### 要求和验收标准

据我们所描述的，我们必须处置汽车和摩托车类型的建筑商，以及一个称为ManufacturingDirector的独特总监来接管建筑商和制造产品。 因此，“车辆制造商”示例的要求如下：

* 我必须拥有能够构造车辆所需的一切的制造类型

* 使用汽车制造商时，必须返回带有四个轮子，五个座椅和定义为`Car`的结构的`VehicleProduct`

* 使用摩托车制造商时，必须返回带有两个轮子，两个座椅和一个定义为`Motorbike`的结构的`VehicleProduct`

* 任何`BuildProcess`构建器构建的`VehicleProduct`必须开放修改

### 车辆制造商的单元测试

根据先前的验收标准，我们将创建一个Director变量，即ManufacturingDirector类型，以使用汽车和摩托车的产品制造商变量代表的构建过程。 导演是负责建造对象的人，而建造者是负责退还实际车辆的人。 因此，我们的生成器声明如下所示：

```go
package creational

type BuildProcess interface {
       SetWheels() BuildProcess
       SetSeats() BuildProcess
       SetStructure() BuildProcess
       GetVehicle() VehicleProduct
}
```

前面的界面定义了制造车辆所需的步骤。 如果要由制造商使用，则每个制造商都必须实现此接口。 在`Set`的每个步骤中，我们都返回相同的构建过程，因此我们可以在同一条语句中将各个步骤链接在一起，这将在后面看到。 最后，我们需要一个`GetVehicle` 方法来从构建器中检索`Vehicle`实例:

```go
type ManufacturingDirector struct {}

func (f *ManufacturingDirector) Construct() {
    //Implementation goes here
}

func（f * ManufacturingDirector）SetBuilder（b BuildProcess）{
  //实现在这里
}
```

`ManufacturingDirector` 是负责接受建造者的变量。 它具有一个`Construct`方法，该方法将使用存储在`Manufacturing`中的构建器，并复制所需的步骤。`SetBuilder`方法将允许我们更改`Manufacturing` director中使用的构建器：

```go
type VehicleProduct struct {
    Wheels    int
    Seats int
    Structure string
}
```

产品是我们在使用制造业时要检索的最终对象。 在这种情况下，车辆由车轮，座椅和结构组成：

```go
type CarBuilder struct {}

func (c *CarBuilder) SetWheels() BuildProcess {
    return nil
}

func (c *CarBuilder) SetSeats() BuildProcess {
    return nil
}

func (c *CarBuilder) SetStructure() BuildProcess {
    return nil
}

func (c *CarBuilder) Build() VehicleProduct {
    return VehicleProduct{}
}
```

第一位建筑商是汽车制造商。 它必须实现BuildProcess接口中定义的每个方法。 在这里，我们将为该特定的构建器设置信息：

```go
type BikeBuilder struct {}

func (b *BikeBuilder) SetWheels() BuildProcess {
    return nil
}

func (b *BikeBuilder) SetSeats() BuildProcess {
    return nil
}

func (b * BikeBuilder) SetStructure() BuildProcess {
  return nil
}

func (b * BikeBuilder) Build() VehicleProduct {
  return VehicleProduct {}
}
```

`Motorbike`结构必须与`Car`结构相同，因为它们都是`Builder`的实现，但请记住，构建每种摩托车的过程可能会非常不同。 使用此对象声明，我们可以创建以下测试：

```go
package creational

import "testing"

func TestBuilderPattern(t *testing.T) {
  manufacturingComplex := ManufacturingDirector{}
  carBuilder := &CarBuilder{}
  manufacturingComplex.SetBuilder(carBuilder)
  manufacturingComplex.Construct()
  car := carBuilder.Build()
  //code continues here...
```

我们将从制造总监和汽车制造商开始，以满足前两个验收标准。 在前面的代码中，我们将创建我们的制造总监，负责测试过程中每辆车的创建。 创建制造主管后，我们创建了CarBuilder，然后使用SetBuilder方法将其传递到制造中。 一旦制造总监知道现在要构造什么，我们就可以调用Construct方法使用CarBuilder创建VehicleProduct。 最后，一旦我们拥有了汽车的所有部件，就可以在CarBuilder上调用GetVehicle方法来检索Car实例：

```go
if car.Wheels != 4 {
  t.Errorf("Wheels on a car must be 4 and they were %d\n", car.Wheels)
}

if car.Structure != "Car" {
  t.Errorf("Structure on a car must be 'Car' and was %s\n",car.Structure)
}

if car.Seats != 5 {
  t.Errorf("Seats on a car must be 5 and they were %d\n", car.Seats)
}
```
我们编写了三个小测试，以检查结果是否为汽车。 我们检查了汽车是否有四个轮子，结构是否为汽车，座位数为五个。 我们有足够的数据来执行测试并确保测试失败，因此我们可以将其视为可靠的:
