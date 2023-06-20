
CREATE TABLE `transaction_detail` (
  `id_transaction` int(11) NOT NULL,
  `id_product` int(11) NOT NULL,
  `quantity` int(11) NOT NULL,
  `bill` int(11) NOT NULL,
  `created_at` datetime NOT NULL DEFAULT current_timestamp()
);

ALTER TABLE `transaction_detail`
  ADD UNIQUE KEY `id_transaction` (`id_transaction`,`id_product`),
  ADD KEY `id_product` (`id_product`);


ALTER TABLE `transaction_detail`
  ADD CONSTRAINT `transaction_detail_ibfk_1` FOREIGN KEY (`id_product`) REFERENCES `product` (`id_product`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  ADD CONSTRAINT `transaction_detail_ibfk_2` FOREIGN KEY (`id_transaction`) REFERENCES `transaction` (`id_transaction`) ON DELETE NO ACTION ON UPDATE NO ACTION;
COMMIT;
