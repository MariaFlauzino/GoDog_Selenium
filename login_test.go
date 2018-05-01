package login

import (
	"regexp"
	"strings"
	"github.com/DATA-DOG/godog/gherkin"
	"fmt"
	"github.com/maria/mark/support"
	"github.com/tebeka/selenium"
	"github.com/DATA-DOG/godog"
)

var Driver selenium.WebDriver


func queEuAcesseiAPginaPrincipal() error {
	Driver.Get("http://opensource.demo.orangehrmlive.com/")
	return nil
}

func faoLoginComE(email, senha string) (err error){
	campoEmail, err := Driver.FindElement(selenium.ByID, "txtUsername")

	if err != nil{
		return
	}
	campoEmail.SendKeys(email)

	campoSenha, err := Driver.FindElement(selenium.ByID, "txtPassword")

	if err != nil{
		return
	}
	campoSenha.SendKeys(senha)

	botaoLogin, err := Driver.FindElement(selenium.ByID,"btnLogin")
	if err != nil{
		return
	}
	botaoLogin.Click()

	return nil
}

func souAutenticadoComSucesso() (err error) {

	nomeMenu, err := Driver.FindElement(selenium.ByClassName,"panelTrigger")
	if err != nil{
		return
	}
	
	saida, _ := nomeMenu.Text()

	if saida != "Welcome Admin"{
		return fmt.Errorf("Erro ao validar usuário atenticado")
	}
	return nil
}

func deveSerASeguinteMensagem(mensagem string) (err error) {

	divAlerta, err := Driver.FindElement(selenium.ByCSSSelector,"#spanMessage")
	
	if err != nil{
		return
	}

	saida, _ := divAlerta.Text()

	if saida != mensagem {
		return fmt.Errorf("Esperado: %v - Obtido: %v", mensagem, saida)
	}
	return nil
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^que eu acessei a página principal$`, queEuAcesseiAPginaPrincipal)
	s.Step(`^faço login com "([^"]*)" e "([^"]*)"$`, faoLoginComE)
	s.Step(`^sou autenticado com sucesso$`, souAutenticadoComSucesso)
	s.Step(`^deve ser a seguinte mensagem "([^"]*)"$`, deveSerASeguinteMensagem)

	s.BeforeScenario(func(interface{}){
		Driver = support.WDInit()
	})

	s.AfterScenario(func(i interface{}, e error){
		
		sc := i.(*gherkin.Scenario)

		rgex := regexp.MustCompile("[^0-9a-zA-Z]+")
		fileName := strings.ToLower(rgex.ReplaceAllString(sc.Name, "-"))
		
		shot,_ := Driver.Screenshot()

		support.SaveImage(shot,fileName)

		fmt.Println(fileName)
		Driver.Quit()


		
	})
}