CREATE TABLE IF NOT EXISTS Orders 
		( Id int NOT NULL PRIMARY KEY AUTO_INCREMENT, 
			UserName varchar(255) DEFAULT NULL, 
			TotalPrice decimal(18,2) NOT NULL, 
			FirstName varchar(255) DEFAULT NULL, 
			LastName varchar(255) DEFAULT NULL, 
			EmailAddress varchar(255) DEFAULT NULL, 
			AddressLine varchar(255) DEFAULT NULL, 
			Country varchar(255) DEFAULT NULL, 
			State varchar(255) DEFAULT NULL, 
			ZipCode varchar(255) DEFAULT NULL, 
			CardName varchar(255) DEFAULT NULL, 
			CardNumber varchar(255) DEFAULT NULL, 
			Expiration varchar(255) DEFAULT NULL, 
			CVV varchar(255) DEFAULT NULL, 
			PaymentMethod int NOT NULL, 
			CreatedBy varchar(255) DEFAULT NULL, 
			CreatedDate datetime NOT NULL, 
			LastModifiedBy varchar(255) 
			DEFAULT NULL, 
			LastModifiedDate datetime DEFAULT NULL
			);