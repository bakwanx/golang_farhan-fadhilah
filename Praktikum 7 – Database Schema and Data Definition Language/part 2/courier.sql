CREATE TABLE `courier` (
  `id_courier` int(11) NOT NULL,
  `courier` varchar(255) NOT NULL
);

ALTER TABLE `courier` ADD PRIMARY KEY (`id_courier`);


ALTER TABLE `courier`
  MODIFY `id_courier` int(11) NOT NULL AUTO_INCREMENT;
COMMIT;

ALTER TABLE `courier` ADD `ongkos_dasar` INT NOT NULL AFTER `courier`;

RENAME TABLE courier To shipping;

DROP TABLE shipping;
