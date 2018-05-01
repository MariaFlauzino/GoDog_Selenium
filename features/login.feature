#language: pt

Funcionalidade: Login
    Para que eu possa ter acesso as minhas tarefas
    Sendo um usuário
    Posso me autenticar com os meus dados previamente cadastrados

    Contexto: Formulário de Login
        Dado que eu acessei a página principal

    Cenario: Login do usuário

        Quando faço login com "admin" e "admin"
        Então sou autenticado com sucesso

    Cenario: Senha incorreta

        Quando faço login com "eu@papito.io" e "xpto123456"
        Então deve ser a seguinte mensagem "Invalid credentials"

    Cenario: E-mail incorreto

        Quando faço login com "eu@papito.io" e "xpto123456"
        Então deve ser a seguinte mensagem "Invalid credentials"