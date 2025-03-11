# Nome do arquivo de cobertura
COVERAGE_FILE=coverage.out
COVERAGE_HTML=coverage.html
COVERAGE_THRESHOLD=80.0 # Defina a cobertura mínima aceitável (exemplo: 80%)

# Gera a cobertura de testes
test:
	@echo "🔍 Executando testes..."
	@go test -coverprofile=$(COVERAGE_FILE) ./...

# Exibe a cobertura no terminal
coverage:
	@echo "📊 Exibindo cobertura de testes..."
	@go tool cover -func=$(COVERAGE_FILE)

# Verifica se a cobertura atinge o mínimo esperado
check-coverage: test coverage
	@COVERAGE=$$(go tool cover -func=$(COVERAGE_FILE) | grep total | awk '{print substr($$3, 1, length($$3)-1)}'); \
	echo "✅ Cobertura atual: $$COVERAGE%"; \
	COVERAGE_NUM=$$(echo $$COVERAGE | awk '{print ($$1+0)}'); \
	if [ $$(echo "$$COVERAGE_NUM < $(COVERAGE_THRESHOLD)" | bc -l) -eq 1 ]; then \
		echo "❌ Cobertura abaixo do esperado ($(COVERAGE_THRESHOLD)%)"; \
		exit 1; \
	else \
		echo "🎉 Cobertura dentro do esperado!"; \
	fi

# Gera um relatório HTML da cobertura
coverage-html: test
	@echo "🌐 Gerando relatório HTML de cobertura..."
	@go tool cover -html=$(COVERAGE_FILE) -o $(COVERAGE_HTML)
	@echo "📂 Abra o arquivo $(COVERAGE_HTML) no navegador."

# Limpa arquivos temporários
clean:
	@echo "🧹 Limpando arquivos de cobertura..."
	@rm -f $(COVERAGE_FILE) $(COVERAGE_HTML)
