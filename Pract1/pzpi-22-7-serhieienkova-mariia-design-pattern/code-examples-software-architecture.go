// Приклад проблеми без Visitor 
type Patient struct {
name string
needsCheckup bool
needsMedication bool
}

func (p *Patient) PerformDailyCare() {
if p.needsCheckup {
// перевірка стану
}
if p.needsMedication {
// нагадування про ліки
}
}

// Рішення проблеми через Visitor
type Visitor interface {
VisitPatient(p *Patient)
}

type Patient struct {
name string
needsCheckup bool
needsMedication bool
}

func (p *Patient) Accept(v Visitor) {
v.VisitPatient(p)
}

// Наведення другого прикладу без патерну Visitor
type ElderlyPatient struct {
Name string
Age  int
}

type DisabledPatient struct {
Name string
Disability string
}

func provideCareForElderly(p ElderlyPatient) {
fmt.Printf("Providing elderly care to %s (age %d)\n", p.Name, p.Age)
}

func provideCareForDisabled(p DisabledPatient) {
fmt.Printf("Providing disability care to %s (%s)\n", p.Name, p.Disability)
}

func main() {
el := ElderlyPatient{"Maria", 80}
ds := DisabledPatient{"Ivan", "Mobility impairment"}

provideCareForElderly(el)
provideCareForDisabled(ds)
}

//Впровадження Visitor
// Visitor
type CareVisitor interface {
VisitElderly(*ElderlyPatient)
VisitDisabled(*DisabledPatient)
}

// Element
type Patient interface {
Accept(CareVisitor)
}

// ConcreteElement 1
type ElderlyPatient struct {
Name string
Age  int
}

func (p *ElderlyPatient) Accept(v CareVisitor) {
v.VisitElderly(p)
}

// ConcreteElement 2
type DisabledPatient struct {
Name       string
Disability string
}

func (p *DisabledPatient) Accept(v CareVisitor) {
v.VisitDisabled(p)
}

// ConcreteVisitor
type DailyCare struct{}

func (c *DailyCare) VisitElderly(p *ElderlyPatient) {
fmt.Printf("Daily care for elderly: %s, %d y.o.\n", p.Name, p.Age)
}

func (c *DailyCare) VisitDisabled(p *DisabledPatient) {
fmt.Printf("Daily care for disabled: %s (%s)\n", p.Name, p.Disability)
}

func main() {
patients := []Patient{
&ElderlyPatient{Name: "Maria", Age: 80},
&DisabledPatient{Name: "Ivan", Disability: "Mobility impairment"},
}

visitor := &DailyCare{}
for _, p := range patients {
p.Accept(visitor)
}
}

