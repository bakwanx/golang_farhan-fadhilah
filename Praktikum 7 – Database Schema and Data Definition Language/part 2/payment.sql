
CREATE TABLE `payment` (
  `id_payment` int(11) NOT NULL,
  `payment` varchar(100) NOT NULL
);

ALTER TABLE `payment`
  ADD PRIMARY KEY (`id_payment`);

ALTER TABLE `payment`
  MODIFY `id_payment` int(11) NOT NULL AUTO_INCREMENT;
COMMIT;
