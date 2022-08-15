package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/Natanael-devops/api-orcamento-familiar/controllers"
	"github.com/Natanael-devops/api-orcamento-familiar/database"
	"github.com/Natanael-devops/api-orcamento-familiar/models"
	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
)

var ID int

func SetupDasRotasDeTeste() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	rotas := gin.Default()
	return rotas
}

//funções para mockar receitas e despesas de teste. Para não correr o risco de efetuar teste com body vazio e dar status 200.
func CriaDespesaMock() {
	descricao := "DespesaMock"

	valor := float32(300.00)
	data := "05/1997"
	categoria := ""
	despesa := models.Despesa{Descricao: descricao, Valor: valor, Data: data, Categoria: categoria}
	database.DB.Create(&despesa)
	ID = int(despesa.ID)
}

func DeletaDespesaMock() {
	var despesa models.Despesa
	database.DB.Delete(&despesa, ID)
}

func CriaReceitaMock() {
	descricao := "ReceitaMock"
	valor := float32(400.00)
	data := "05/1997"
	receita := models.Receita{Descricao: descricao, Valor: valor, Data: data}
	database.DB.Create(&receita)
	ID = int(receita.ID)

}

func DeletaReceitaMock() {
	var receita models.Receita
	database.DB.Delete(&receita, ID)
}

//próximas funções testam método GET da listagem de todas as receitas e despesas.
func TestListagemDasReceitas(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaReceitaMock()
	defer DeletaReceitaMock()
	r := SetupDasRotasDeTeste()
	r.GET("/receitas", controllers.ExibeTodasReceitas)
	req, _ := http.NewRequest("GET", "/receitas", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
	fmt.Println(resposta.Body)
}

func TestListagemDasDespesas(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaDespesaMock()
	defer DeletaDespesaMock()
	r := SetupDasRotasDeTeste()
	r.GET("/despesas", controllers.ExibeTodasDespesas)
	req, _ := http.NewRequest("GET", "/despesas", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
	fmt.Println(resposta.Body)
}

//próximas funções testam a busca por ID das receitas e das despesas.
func TestBuscaReceitaPorID(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaReceitaMock()
	defer DeletaReceitaMock()
	r := SetupDasRotasDeTeste()
	r.GET("/receitas/:id", controllers.DetalhaReceitaPorID)
	path := "/receitas/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("GET", path, nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestBuscaDespesaPorID(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaDespesaMock()
	defer DeletaDespesaMock()
	r := SetupDasRotasDeTeste()
	r.GET("/despesas/:id", controllers.DetalhaDespesaPorID)
	path := "/despesas/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("GET", path, nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}

//próximas funções testam o método DELETE das receitas e despesas

func TestDeletaReceita(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaReceitaMock()
	r := SetupDasRotasDeTeste()
	r.DELETE("/receitas/:id", controllers.DeletaReceita)
	path := "/receitas/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("DELETE", path, nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestDeletaDespesa(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaDespesaMock()
	r := SetupDasRotasDeTeste()
	r.DELETE("/despesas/:id", controllers.DeletaDespesa)
	path := "/despesas/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("DELETE", path, nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}

//próximas funções testam o método PUT das receitas e despesas.

func TestEditaReceita(t *testing.T) {
	descricao := "OutraReceitaMock"
	valor := float32(500.00)
	data := "dezembro"
	database.ConectaComBancoDeDados()
	CriaReceitaMock()
	defer DeletaReceitaMock()
	r := SetupDasRotasDeTeste()
	r.PUT("/receitas/:id", controllers.EditaReceita)
	receita := models.Receita{Descricao: descricao, Valor: valor, Data: data}
	valorJson, _ := json.Marshal(receita)
	path := "/receitas/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("PUT", path, bytes.NewBuffer(valorJson))
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	var receitaMockAtualizada models.Receita
	json.Unmarshal(resposta.Body.Bytes(), &receitaMockAtualizada)
	assert.Equal(t, "OutraReceitaMock", receitaMockAtualizada.Descricao)
	assert.Equal(t, "dezembro", receitaMockAtualizada.Data)
	assert.Equal(t, float32(500.00), receitaMockAtualizada.Valor)
}

func TestEditaDespesa(t *testing.T) {
	descricao := "OutraDespesaMock"
	valor := float32(700.00)
	data := "janeiro"
	database.ConectaComBancoDeDados()
	CriaReceitaMock()
	defer DeletaReceitaMock()
	r := SetupDasRotasDeTeste()
	r.PUT("/despesas/:id", controllers.EditaDespesa)
	despesa := models.Despesa{Descricao: descricao, Valor: valor, Data: data}
	valorJson, _ := json.Marshal(despesa)
	path := "/despesas/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("PUT", path, bytes.NewBuffer(valorJson))
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	var despesaMockAtualizada models.Despesa
	json.Unmarshal(resposta.Body.Bytes(), &despesaMockAtualizada)
	assert.Equal(t, "OutraDespesaMock", despesaMockAtualizada.Descricao)
	assert.Equal(t, "janeiro", despesaMockAtualizada.Data)
	assert.Equal(t, float32(700.00), despesaMockAtualizada.Valor)
}

func TestBuscaReceitasPorMes(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaReceitaMock()
	r := SetupDasRotasDeTeste()
	defer DeletaReceitaMock()

	r.GET("/receitas/:id/:mes", controllers.BuscaReceitaPorMes)
	path := "/receitas/2022/08"
	req, _ := http.NewRequest("GET", path, nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestBuscaDespesasPorMes(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaDespesaMock()
	r := SetupDasRotasDeTeste()
	defer DeletaDespesaMock()

	r.GET("/despesas/:id/:mes", controllers.BuscaReceitaPorMes)
	path := "/despesas/2022/08"
	req, _ := http.NewRequest("GET", path, nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestBuscaReceitasPorDescricao(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaReceitaMock()
	r := SetupDasRotasDeTeste()
	defer DeletaReceitaMock()
	path := "/receitas?descricao="
	path += "ReceitaMock"
	r.GET("/receitas", controllers.BuscaReceitaPorDescricao)
	req, _ := http.NewRequest("GET", path, nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)

}

func TestBuscaDespesasPorDescricao(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaDespesaMock()
	r := SetupDasRotasDeTeste()
	defer DeletaDespesaMock()
	path := "/despesas?descricao="
	path += "DespesaMock"
	r.GET("/despesas", controllers.BuscaDespesaPorDescricao)
	req, _ := http.NewRequest("GET", path, nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestResumoMensal(t *testing.T) {
	database.ConectaComBancoDeDados()
	CriaReceitaMock()
	CriaDespesaMock()
	defer DeletaReceitaMock()
	defer DeletaDespesaMock()
	r := SetupDasRotasDeTeste()
	path := "/resumo/1997/05"
	r.GET("/resumo/:id/:mes", controllers.DevolveResumoMensal)
	req, _ := http.NewRequest("GET", path, nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)

}
