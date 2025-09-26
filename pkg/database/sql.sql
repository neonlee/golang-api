-- Tabela de Planos/Assinaturas
CREATE TABLE planos (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(50) NOT NULL, -- Basic, Premium, Master
    descricao TEXT,
    valor_mensal DECIMAL(10,2),
    modulos_disponiveis JSONB, -- Array com módulos permitidos
    limite_usuarios INT,
    limite_empresas INT,
    ativo BOOLEAN DEFAULT true,
    created_at TIMESTAMP DEFAULT NOW()
);

-- Tabela de Empresas/Clientess
CREATE TABLE empresas (
    id SERIAL PRIMARY KEY,
    nome_fantasia VARCHAR(100) NOT NULL,
    razao_social VARCHAR(100),
    cnpj VARCHAR(18) UNIQUE,
    telefone VARCHAR(15),
    email VARCHAR(100),
    endereco JSONB,
    plano_id INT REFERENCES planos(id),
    data_assinatura DATE,
    status VARCHAR(20) DEFAULT 'ativo',
    created_at TIMESTAMP DEFAULT NOW()
);

-- Tabela de Módulos do Sistema
CREATE TABLE modulos (
    id SERIAL PRIMARY KEY,
    nome VARCHAR(50) NOT NULL,
    descricao TEXT,
    categoria VARCHAR(30), -- petshop, veterinario, hotel, racao
    icone VARCHAR(30),
    ordem INT,
    ativo BOOLEAN DEFAULT true
);

-- Tabela de Usuários
CREATE TABLE usuarios (
    id SERIAL PRIMARY KEY,
    empresa_id INT REFERENCES empresas(id),
    nome VARCHAR(100) NOT NULL,
    email VARCHAR(100) UNIQUE NOT NULL,
    senha_hash VARCHAR(255) NOT NULL,
    telefone VARCHAR(15),
    cargo VARCHAR(50),
    foto_url VARCHAR(255),
    ativo BOOLEAN DEFAULT true,
    ultimo_login TIMESTAMP,
    created_at TIMESTAMP DEFAULT NOW()
);

-- Tabela de Perfis de Acesso
CREATE TABLE perfis (
    id SERIAL PRIMARY KEY,
    empresa_id INT REFERENCES empresas(id),
    nome VARCHAR(50) NOT NULL,
    descricao TEXT,
    is_admin BOOLEAN DEFAULT false,
    created_at TIMESTAMP DEFAULT NOW()
);

-- Tabela de Permissões por Módulo
CREATE TABLE permissoes (
    id SERIAL PRIMARY KEY,
    perfil_id INT REFERENCES perfis(id),
    modulo_id INT REFERENCES modulos(id),
    pode_visualizar BOOLEAN DEFAULT false,
    pode_editar BOOLEAN DEFAULT false,
    pode_excluir BOOLEAN DEFAULT false,
    pode_gerar_relatorio BOOLEAN DEFAULT false
);

-- Tabela de Usuário x Perfil
CREATE TABLE usuario_perfis (
    usuario_id INT REFERENCES usuarios(id),
    perfil_id INT REFERENCES perfis(id),
    PRIMARY KEY (usuario_id, perfil_id)
);

-- Tabela de Clientess
CREATE TABLE clientes (
    id SERIAL PRIMARY KEY,
    empresa_id INT REFERENCES empresas(id),
    nome VARCHAR(100) NOT NULL,
    cpf_cnpj VARCHAR(18),
    telefone VARCHAR(15),
    email VARCHAR(100),
    endereco JSONB,
    data_nascimento DATE,
    observacoes TEXT,
    ativo BOOLEAN DEFAULT true,
    created_at TIMESTAMP DEFAULT NOW()
);

-- Tabela de Pets
CREATE TABLE pets (
    id SERIAL PRIMARY KEY,
    cliente_id INT REFERENCES clientes(id),
    nome VARCHAR(50) NOT NULL,
    especie VARCHAR(20), -- Cachorro, Gato, etc.
    raca VARCHAR(50),
    sexo VARCHAR(1), -- M, F
    data_nascimento DATE,
    peso DECIMAL(5,2),
    cor VARCHAR(30),
    observacoes TEXT,
    foto_url VARCHAR(255),
    created_at TIMESTAMP DEFAULT NOW()
);

-- Tabela de Fornecedores
CREATE TABLE fornecedores (
    id SERIAL PRIMARY KEY,
    empresa_id INT REFERENCES empresas(id),
    nome_fantasia VARCHAR(100) NOT NULL,
    razao_social VARCHAR(100),
    cnpj VARCHAR(18),
    telefone VARCHAR(15),
    email VARCHAR(100),
    endereco JSONB,
    ativo BOOLEAN DEFAULT true,
    created_at TIMESTAMP DEFAULT NOW()
);

-- Tabela de Categorias de Produtos
CREATE TABLE categorias_produtos (
    id SERIAL PRIMARY KEY,
    empresa_id INT REFERENCES empresas(id),
    nome VARCHAR(50) NOT NULL,
    descricao TEXT,
    ativo BOOLEAN DEFAULT true
);

-- Tabela de Produtos
CREATE TABLE produtos (
    id SERIAL PRIMARY KEY,
    empresa_id INT REFERENCES empresas(id),
    categoria_id INT REFERENCES categorias_produtos(id),
    fornecedor_id INT REFERENCES fornecedores(id),
    codigo_barras VARCHAR(50),
    nome VARCHAR(100) NOT NULL,
    descricao TEXT,
    tipo VARCHAR(20), -- racao, medicamento, brinquedo, etc.
    especie_destinada VARCHAR(20), -- Cachorro, Gato, etc.
    peso_kg DECIMAL(8,3),
    unidade_medida VARCHAR(10),
    preco_custo DECIMAL(10,2),
    preco_venda DECIMAL(10,2),
    estoque_minimo INT,
    ativo BOOLEAN DEFAULT true,
    created_at TIMESTAMP DEFAULT NOW()
);

-- Tabela de Movimentação de Estoque
CREATE TABLE movimentacao_estoque (
    id SERIAL PRIMARY KEY,
    produto_id INT REFERENCES produtos(id),
    tipo_movimentacao VARCHAR(20), -- entrada, saida, ajuste
    quantidade INT NOT NULL,
    quantidade_anterior INT,
    quantidade_atual INT,
    motivo VARCHAR(100),
    usuario_id INT REFERENCES usuarios(id),
    created_at TIMESTAMP DEFAULT NOW()
);

-- Tabela de Tipos de Serviços
CREATE TABLE tipos_servicos (
    id SERIAL PRIMARY KEY,
    empresa_id INT REFERENCES empresas(id),
    nome VARCHAR(50) NOT NULL,
    categoria VARCHAR(30), -- banho, tosa, consulta, hospedagem
    duracao_minutos INT,
    valor DECIMAL(10,2),
    descricao TEXT,
    ativo BOOLEAN DEFAULT true
);

-- Tabela de Agendamentos
CREATE TABLE agendamentos (
    id SERIAL PRIMARY KEY,
    empresa_id INT REFERENCES empresas(id),
    cliente_id INT REFERENCES clientes(id),
    pet_id INT REFERENCES pets(id),
    tipo_servico_id INT REFERENCES tipos_servicos(id),
    data_agendamento TIMESTAMP NOT NULL,
    status VARCHAR(20), -- agendado, confirmado, em_andamento, concluido, cancelado
    observacoes TEXT,
    valor_estimado DECIMAL(10,2),
    usuario_id INT REFERENCES usuarios(id),
    created_at TIMESTAMP DEFAULT NOW()
);

-- Tabela de Prontuários Veterinários
CREATE TABLE prontuarios (
    id SERIAL PRIMARY KEY,
    pet_id INT REFERENCES pets(id),
    veterinario_id INT REFERENCES usuarios(id),
    data_consulta TIMESTAMP DEFAULT NOW(),
    anamnese TEXT,
    diagnostico TEXT,
    prescricao JSONB, -- Array de medicamentos
    observacoes TEXT,
    peso_atual DECIMAL(5,2),
    temperatura DECIMAL(4,2),
    created_at TIMESTAMP DEFAULT NOW()
);

-- Tabela de Vendas
CREATE TABLE vendas (
    id SERIAL PRIMARY KEY,
    empresa_id INT REFERENCES empresas(id),
    cliente_id INT REFERENCES clientes(id),
    usuario_id INT REFERENCES usuarios(id),
    data_venda TIMESTAMP DEFAULT NOW(),
    tipo_venda VARCHAR(20), -- balcao, delivery, online
    status VARCHAR(20), -- pendente, pago, cancelado
    valor_total DECIMAL(10,2),
    desconto DECIMAL(10,2),
    valor_final DECIMAL(10,2),
    forma_pagamento VARCHAR(30),
    observacoes TEXT
);

-- Tabela de Itens da Venda
CREATE TABLE venda_itens (
    id SERIAL PRIMARY KEY,
    venda_id INT REFERENCES vendas(id),
    produto_id INT REFERENCES produtos(id),
    tipo_servico_id INT REFERENCES tipos_servicos(id),
    quantidade INT,
    valor_unitario DECIMAL(10,2),
    valor_total DECIMAL(10,2),
    tipo_item VARCHAR(10) -- produto ou servico
);

-- Tabela de Contas a Receber
CREATE TABLE contas_receber (
    id SERIAL PRIMARY KEY,
    venda_id INT REFERENCES vendas(id),
    cliente_id INT REFERENCES clientes(id),
    descricao VARCHAR(100),
    valor DECIMAL(10,2),
    data_vencimento DATE,
    data_pagamento DATE,
    status VARCHAR(20), -- pendente, pago, atrasado
    forma_pagamento VARCHAR(30)
);

-- Tabela de Configuração NF-e
CREATE TABLE nfe_config (
    id SERIAL PRIMARY KEY,
    empresa_id INT REFERENCES empresas(id) UNIQUE,
    ambiente VARCHAR(10) DEFAULT 'homologacao', -- homologacao, producao
    certificado_digital BYTEA,
    senha_certificado VARCHAR(100),
    sequencia_nfe INT DEFAULT 1,
    serie_nfe INT DEFAULT 1,
    token_csc VARCHAR(100),
    csc VARCHAR(100),
    created_at TIMESTAMP DEFAULT NOW()
);

-- Tabela de NF-e
CREATE TABLE nfe (
    id SERIAL PRIMARY KEY,
    empresa_id INT REFERENCES empresas(id),
    venda_id INT REFERENCES vendas(id),
    numero_nfe INT,
    serie_nfe INT,
    chave_acesso VARCHAR(44),
    status VARCHAR(20), -- pendente, autorizada, cancelada, erro
    xml_nfe TEXT,
    protocolo VARCHAR(50),
    data_emissao TIMESTAMP,
    data_autorizacao TIMESTAMP,
    motivo_cancelamento TEXT,
    created_at TIMESTAMP DEFAULT NOW()
);

-- Tabela de Categorias de Despesas
CREATE TABLE categorias_despesas (
    id SERIAL PRIMARY KEY,
    empresa_id INT REFERENCES empresas(id),
    nome VARCHAR(50) NOT NULL,
    descricao TEXT,
    tipo VARCHAR(20), -- fixa, variavel, ocasional
    ativo BOOLEAN DEFAULT true
);

-- Tabela de Contas a Pagar
CREATE TABLE contas_pagar (
    id SERIAL PRIMARY KEY,
    empresa_id INT REFERENCES empresas(id),
    fornecedor_id INT REFERENCES fornecedores(id),
    categoria_despesa_id INT REFERENCES categorias_despesas(id),
    descricao VARCHAR(100) NOT NULL,
    numero_documento VARCHAR(50),
    valor_original DECIMAL(10,2),
    valor_juros DECIMAL(10,2) DEFAULT 0,
    valor_multa DECIMAL(10,2) DEFAULT 0,
    valor_desconto DECIMAL(10,2) DEFAULT 0,
    valor_final DECIMAL(10,2),
    data_emissao DATE,
    data_vencimento DATE,
    data_pagamento DATE,
    status VARCHAR(20) DEFAULT 'pendente', -- pendente, pago, atrasado, cancelado
    forma_pagamento VARCHAR(30),
    observacoes TEXT,
    usuario_id INT REFERENCES usuarios(id),
    created_at TIMESTAMP DEFAULT NOW()
);

-- Tabela de Parcelas de Contas a Pagar
CREATE TABLE contas_pagar_parcelas (
    id SERIAL PRIMARY KEY,
    conta_pagar_id INT REFERENCES contas_pagar(id),
    numero_parcela INT,
    valor_parcela DECIMAL(10,2),
    data_vencimento DATE,
    data_pagamento DATE,
    status VARCHAR(20) DEFAULT 'pendente',
    juros_calculado DECIMAL(10,2) DEFAULT 0,
    multa_calculada DECIMAL(10,2) DEFAULT 0
);

-- Tabela de Compras/Entradas de Mercadoria
CREATE TABLE compras (
    id SERIAL PRIMARY KEY,
    empresa_id INT REFERENCES empresas(id),
    fornecedor_id INT REFERENCES fornecedores(id),
    numero_nota_fiscal VARCHAR(50),
    data_compra DATE,
    data_entrada DATE,
    valor_total DECIMAL(10,2),
    valor_frete DECIMAL(10,2) DEFAULT 0,
    valor_desconto DECIMAL(10,2) DEFAULT 0,
    status VARCHAR(20) DEFAULT 'pendente', -- pendente, finalizada, cancelada
    observacoes TEXT,
    usuario_id INT REFERENCES usuarios(id),
    created_at TIMESTAMP DEFAULT NOW()
);

-- Tabela de Itens da Compra
CREATE TABLE compra_itens (
    id SERIAL PRIMARY KEY,
    compra_id INT REFERENCES compras(id),
    produto_id INT REFERENCES produtos(id),
    quantidade INT,
    valor_unitario DECIMAL(10,2),
    valor_total DECIMAL(10,2),
    created_at TIMESTAMP DEFAULT NOW()
);

-- Tabela de Contas Bancárias
CREATE TABLE contas_bancarias (
    id SERIAL PRIMARY KEY,
    empresa_id INT REFERENCES empresas(id),
    nome VARCHAR(50) NOT NULL,
    banco VARCHAR(50),
    agencia VARCHAR(10),
    conta VARCHAR(15),
    saldo_inicial DECIMAL(10,2) DEFAULT 0,
    tipo VARCHAR(20), -- corrente, poupanca
    ativo BOOLEAN DEFAULT true,
    created_at TIMESTAMP DEFAULT NOW()
);

-- Tabela de Movimentações Bancárias
CREATE TABLE movimentacoes_bancarias (
    id SERIAL PRIMARY KEY,
    conta_bancaria_id INT REFERENCES contas_bancarias(id),
    tipo_movimentacao VARCHAR(20), -- entrada, saida, transferencia
    valor DECIMAL(10,2),
    data_movimentacao DATE,
    descricao VARCHAR(100),
    categoria VARCHAR(50),
    contas_receber_id INT REFERENCES contas_receber(id),
    contas_pagar_id INT REFERENCES contas_pagar(id),
    observacoes TEXT,
    usuario_id INT REFERENCES usuarios(id),
    created_at TIMESTAMP DEFAULT NOW()
);

-- Tabela de Transferências entre Contas
CREATE TABLE transferencias_bancarias (
    id SERIAL PRIMARY KEY,
    empresa_id INT REFERENCES empresas(id),
    conta_origem_id INT REFERENCES contas_bancarias(id),
    conta_destino_id INT REFERENCES contas_bancarias(id),
    valor DECIMAL(10,2),
    data_transferencia DATE,
    taxa DECIMAL(10,2) DEFAULT 0,
    observacoes TEXT,
    usuario_id INT REFERENCES usuarios(id),
    created_at TIMESTAMP DEFAULT NOW()
);

-- Tabela de Fechamentos de Caixa
CREATE TABLE fechamentos_caixa (
    id SERIAL PRIMARY KEY,
    empresa_id INT REFERENCES empresas(id),
    usuario_id INT REFERENCES usuarios(id),
    data_fechamento DATE,
    valor_inicial DECIMAL(10,2),
    valor_final DECIMAL(10,2),
    valor_vendas DECIMAL(10,2),
    valor_retiradas DECIMAL(10,2),
    valor_diferenca DECIMAL(10,2),
    observacoes TEXT,
    status VARCHAR(20) DEFAULT 'aberto', -- aberto, fechado
    created_at TIMESTAMP DEFAULT NOW()
);

-- Tabela de Demonstrativos Financeiros
CREATE TABLE demonstrativos_financeiros (
    id SERIAL PRIMARY KEY,
    empresa_id INT REFERENCES empresas(id),
    mes_ano VARCHAR(7), -- formato: YYYY-MM
    total_receitas DECIMAL(10,2) DEFAULT 0,
    total_despesas DECIMAL(10,2) DEFAULT 0,
    saldo_final DECIMAL(10,2),
    criado_em TIMESTAMP DEFAULT NOW()
);


-- Tabela de Log do Sistema
CREATE TABLE logs_sistema (
    id SERIAL PRIMARY KEY,
    empresa_id INT REFERENCES empresas(id) ON DELETE SET NULL,
    usuario_id INT REFERENCES usuarios(id) ON DELETE SET NULL,
    modulo VARCHAR(50) NOT NULL,           -- Módulo onde ocorreu a ação
    acao VARCHAR(50) NOT NULL,             -- Ação realizada (INSERT, UPDATE, DELETE, LOGIN, etc.)
    descricao TEXT,                        -- Descrição detalhada da ação
    tabela_afetada VARCHAR(50),            -- Tabela que foi afetada
    registro_id INT,                       -- ID do registro afetado (se aplicável)
    dados_anteriores JSONB,                -- Dados antes da alteração (para UPDATE)
    dados_novos JSONB,                     -- Dados após a alteração (para UPDATE)
    ip_cliente VARCHAR(45),                -- IP do cliente
    user_agent TEXT,                       -- User Agent do navegador
    nivel_log VARCHAR(20) DEFAULT 'INFO',  -- INFO, WARNING, ERROR, DEBUG
    created_at TIMESTAMP DEFAULT NOW()
);

-- Índices para melhor performance
CREATE INDEX idx_logs_empresa ON logs_sistema(empresa_id);
CREATE INDEX idx_logs_usuario ON logs_sistema(usuario_id);
CREATE INDEX idx_logs_modulo ON logs_sistema(modulo);
CREATE INDEX idx_logs_acao ON logs_sistema(acao);
CREATE INDEX idx_logs_data ON logs_sistema(created_at);
CREATE INDEX idx_logs_nivel ON logs_sistema(nivel_log);