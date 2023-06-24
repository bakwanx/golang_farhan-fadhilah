

CREATE TABLE `type` (
  `id_type` int(11) NOT NULL,
  `type` varchar(255) NOT NULL
);

ALTER TABLE `type`
  ADD PRIMARY KEY (`id_type`);


ALTER TABLE `type`
  MODIFY `id_type` int(11) NOT NULL AUTO_INCREMENT;

ALTER TABLE `type`
  ADD CONSTRAINT `type_ibfk_1` FOREIGN KEY (`id_type`) REFERENCES `product` (`id_type`) ON DELETE NO ACTION ON UPDATE NO ACTION;
COMMIT;

