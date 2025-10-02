package controllers

import (
	"fmt"
	"net/http"
	"petApi/internal/models"
	"petApi/internal/repositories"
	"petApi/internal/requests"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	repository repositories.ProdutoRepository
}

func NewProductController(repository repositories.ProdutoRepository) *ProductController {
	return &ProductController{repository: repository}
}

// func (pc *ProductController) GetAllProducts(c *gin.Context) {
// 	products, err := pc.repository.GetAll()
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
// 		return
// 	}
// 	c.JSON(http.StatusOK, products)
// }

func (pc *ProductController) GetProductByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}
	product, err := pc.repository.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}

func (pc *ProductController) CreateProduct(c *gin.Context) {
	var product models.Produtos
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid input: %s", err.Error())})
		return
	}
	err := pc.repository.Create(&product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Failed to create product %s", err.Error())})
		return
	}
	c.JSON(http.StatusCreated, &product)
}

func (pc *ProductController) UpdateProduct(c *gin.Context) {
	var product models.Produtos
	if err := c.ShouldBindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	err := pc.repository.Update(&product)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
		return
	}
	c.JSON(http.StatusOK, &product)
}

func (pc *ProductController) DeleteProduct(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}
	if err := pc.repository.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete product"})
		return
	}
	c.Status(http.StatusNoContent)
}

func (pc *ProductController) ListByEmpresa(c *gin.Context) {
	empresaIDStr := c.Query("empresa_id")
	if empresaIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "empresa_id is required"})
		return
	}
	empresaID, err := strconv.Atoi(empresaIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid empresa_id"})
		return
	}
	var filters requests.ProdutoFilter
	if nome := c.Query("nome"); nome != "" {
		filters.Nome = nome
	}
	if categoriaIDStr := c.Query("categoria_id"); categoriaIDStr != "" {
		categoriaID, err := strconv.Atoi(categoriaIDStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid categoria_id"})
			return
		}
		catIDUint := uint(categoriaID)
		filters.CategoriaID = &catIDUint
	}
	if ativoStr := c.Query("ativo"); ativoStr != "" {
		ativo, err := strconv.ParseBool(ativoStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ativo value"})
			return
		}
		filters.Ativo = &ativo
	}
	if especie := c.Query("especie_destinada"); especie != "" {
		filters.EspecieDestinada = especie
	}
	products, err := pc.repository.ListByEmpresa(uint(empresaID), filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
		return
	}
	c.JSON(http.StatusOK, products)
}

func (pc *ProductController) GetByCategoria(c *gin.Context) {
	categoriaIDStr := c.Param("categoria_id")
	categoriaID, err := strconv.Atoi(categoriaIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid categoria_id"})
		return
	}
	products, err := pc.repository.GetByCategoria(uint(categoriaID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
		return
	}

	c.JSON(http.StatusOK, products)
}

func (pc *ProductController) SearchProducts(c *gin.Context) {
	empresaId := c.Param("empresa_id")
	empresaID, err := strconv.Atoi(empresaId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid categoria_id"})
		return
	}
	termo := c.Query("termo")
	products, err := pc.repository.Search(uint(empresaID), termo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
		return
	}
	c.JSON(http.StatusOK, products)
}

func (pc *ProductController) GetProdutosBaixoEstoque(c *gin.Context) {
	empresaId := c.Param("empresa_id")
	empresaID, err := strconv.Atoi(empresaId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid empresa_id"})
		return
	}
	products, err := pc.repository.GetProdutosBaixoEstoque(uint(empresaID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
		return
	}
	c.JSON(http.StatusOK, products)
}

func (pc *ProductController) GetProdutosVencidos(c *gin.Context) {
	empresaId := c.Param("empresa_id")
	empresaID, err := strconv.Atoi(empresaId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid empresa_id"})
		return
	}
	products, err := pc.repository.GetProdutosVencidos(uint(empresaID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
		return
	}
	c.JSON(http.StatusOK, products)
}

func (pc *ProductController) GetProdutosProximosVencimento(c *gin.Context) {
	empresaId := c.Param("empresa_id")
	empresaID, err := strconv.Atoi(empresaId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid empresa_id"})
		return
	}
	products, err := pc.repository.GetProdutosProximosVencimento(uint(empresaID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
		return
	}
	c.JSON(http.StatusOK, products)
}

func (pc *ProductController) GetProdutosVencimentoHoje(c *gin.Context) {
	empresaId := c.Param("empresa_id")
	empresaID, err := strconv.Atoi(empresaId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid empresa_id"})
		return
	}
	products, err := pc.repository.GetProdutosVencimentoHoje(uint(empresaID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
		return
	}
	c.JSON(http.StatusOK, products)
}

func (pc *ProductController) GetProdutosSemEstoque(c *gin.Context) {
	empresaId := c.Param("empresa_id")
	empresaID, err := strconv.Atoi(empresaId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid empresa_id"})
		return
	}
	products, err := pc.repository.GetProdutosSemEstoque(uint(empresaID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch products"})
		return
	}
	c.JSON(http.StatusOK, products)
}

func (pc *ProductController) GetProdutoComEstoque(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}
	product, err := pc.repository.GetProdutoComEstoque(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
		return
	}
	c.JSON(http.StatusOK, product)
}

func (pc *ProductController) UpdateEstoque(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}
	var req struct {
		Quantidade int `json:"quantidade"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	err = pc.repository.UpdateEstoque(uint(id), req.Quantidade)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update stock"})
		return
	}
	c.Status(http.StatusNoContent)
}
