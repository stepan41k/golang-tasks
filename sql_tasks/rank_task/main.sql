-- Дано: Те же таблицы
-- products (id, name, category_id, price).
-- categories (id, name)

-- Задание:
-- Для каждого товара выведи:
-- 1) Его название (name).
-- 2) Его цену (price).
-- 3) Название категории (category_name).
-- 4) Ранг товара по цене внутри его категории (самый дорогой товар в категории получает ранг 1, следующий — 2 и т.д.).

-- Важный нюанс: Если у двух товаров в одной категории одинаковая цена, они должны получить одинаковый ранг, а следующий за ними товар должен получить ранг через "пропуск" (например: 1, 2, 2, 4).

SELECT p.name, p.price, c.name AS category_name, RANK() OVER (PARTITION BY p.category_id ORDER BY p.price DESC) AS rank
FROM products p
INNER JOIN categories c ON p.category_id = c.id;