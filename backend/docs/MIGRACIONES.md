# Documentación de Migraciones - ANNAK

## Fase 1: Sistema Base

### Users
- UUID como clave primaria
- Campos únicos: username, email
- Tipos enumerados: account_type, gender, preferred_contact, status
- Campos obligatorios: username, email, password_hash, account_type, gender
- Soft delete implementado con deleted_at

### User2FA
- Relacionada con Users (foreign key: user_id)
- Almacena códigos de respaldo en formato JSON
- Flag booleano para indicar si 2FA está activado

### LoginHistory
- Registra historial de inicios de sesión
- Almacena información geográfica y de red
- Relacionada con Users (foreign key: user_id)

## Fase 2: Perfiles Creadores

### CreatorProfile
- Perfil extendido para creadores
- Almacena información de redes sociales en JSON
- Relacionada con Users (foreign key: user_id)

### CreatorLimitations
- Define restricciones de contenido
- Todos los campos son booleanos
- Relacionada con CreatorProfile (foreign key: creator_id)

### CreatorMaterials
- Define tipos de contenido permitidos
- Campos booleanos para cada tipo de material
- Relacionada con CreatorProfile (foreign key: creator_id)

### CreatorVerification
- Almacena documentos de verificación
- Registro de fecha de verificación
- Relacionada con CreatorProfile (foreign key: creator_id)

### CreatorStylePreference
- Preferencias de estilo del creador
- Nivel de comodidad en escala numérica
- Relacionada con CreatorProfile (foreign key: creator_id)

## Fase 3: Taxonomía

### Category
- Sistema jerárquico de categorías
- Campos únicos: slug
- Autorreferencial para categorías padre/hijo

### Keyword
- Sistema de etiquetado simple
- Campos únicos: slug
- Timestamp de creación

### CreatorCategory y CreatorKeyword
- Tablas de relación muchos a muchos
- Relacionan creadores con categorías y keywords

## Fase 4: Producciones

### Production
- Sistema central de contenido
- Campos únicos: slug
- Tipo enumerado: status
- Soft delete implementado
- Arrays para keywords y categorías

### ProductionContent
- Almacena archivos de contenido
- Tipo enumerado: file_type
- Control de contenido AI generado
- Soft delete implementado

### ProductionModeration
- Registro de acciones de moderación
- Relacionada con Productions y Users (admin)

## Fase 5: Interacciones

### Tip
- Sistema de propinas entre usuarios
- Tipo enumerado: status
- Campos decimales para montos

### Subscription
- Sistema de suscripciones
- Tipo enumerado: status
- Control de fechas de inicio/fin

### Interaction
- Sistema genérico de interacciones
- Tipos enumerados: type, content_type
- Polimórfico para diferentes tipos de contenido

### Comment
- Sistema de comentarios
- Tipo enumerado: status
- Soft delete implementado

## Fase 6: Sistema Clientes

### ClientPreferences
- Preferencias de usuario
- Arrays para categorías y keywords preferidas
- Tipo enumerado: nudity_comfort_level

### ClientViewingHistory
- Historial de visualización
- Control de duración y completitud
- Relacionada con Productions

### ClientSearchHistory
- Historial de búsquedas
- Almacena filtros en JSON
- Métricas de resultados

### ClientRecommendation
- Sistema de recomendaciones
- Score decimal para relevancia
- Relacionada con Productions y Creators

## Fase 7: Pagos

### Transaction
- Registro central de transacciones
- Tipos enumerados: type, status
- Campos decimales para montos
- Registro de proveedores de pago

### PaymentMethod
- Métodos de pago de usuarios
- Tokenización de tarjetas
- Flag para método predeterminado

### PayoutAccount
- Cuentas para retiros
- Información de cuenta en JSON
- Sistema de verificación

## Fase 8: Moderación

### ContentReport
- Sistema de reportes de contenido
- Tipos enumerados: status, content_type
- Polimórfico para diferentes tipos de contenido

### ModerationLog
- Registro de acciones de moderación
- Tipos enumerados: action, content_type
- Trazabilidad de decisiones

## Notas Generales

### Campos Comunes
- Todas las tablas usan UUIDs como claves primarias
- Timestamps automáticos (created_at, updated_at)
- Foreign keys con delete cascade donde apropiado
- Soft delete en tablas críticas

### Tipos Enumerados
- Implementados como tipos personalizados en PostgreSQL
- Valores predefinidos para consistencia
- Documentación en constantes de Go

### Índices
- Claves primarias: UUID
- Campos únicos: username, email, slugs
- Foreign keys para optimización de joins
- Índices adicionales según patrones de acceso

### Seguridad
- Contraseñas hasheadas
- Tokenización de datos sensibles
- Separación de datos de verificación
- Logs de auditoría para cambios críticos