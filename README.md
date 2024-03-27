Kullanıcı Ekleme (AddUser): Kullanıcı adı ve parola alanlarını içeren bir HTTP isteği alır. Gelen kullanıcı adı ve parolayı doğrular, ardından parolayı bir hash'e dönüştürür ve bu kullanıcı için bir token oluşturur. Oluşturulan kullanıcıyı veritabanına ekler ve kullanıcıya oluşturulan token ile birlikte kullanıcı bilgilerini döndürür.

Kullanıcı Kaldırma (RemoveUser): Belirli bir kullanıcı adını alır ve bu kullanıcıyı veritabanından siler.

Kullanıcıyı Etkinleştirme (ActivateUser): Belirli bir kullanıcı adını alır, bu kullanıcıyı veritabanında bulur ve kullanıcının token'ını günceller. Daha sonra güncellenmiş kullanıcı bilgilerini döndürür.

Kullanıcıyı Devre Dışı Bırakma (DeactivateUser): Belirli bir kullanıcı adını alır, bu kullanıcıyı veritabanında bulur ve kullanıcının token'ını kaldırır.
