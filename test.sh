#curl -d "@p.xml" -X POST http://localhost:1492/sifxml/Fred
echo '\ncurl -d "@p.xml" -X POST http://localhost:1492/sifxml/PurchaseOrders\n\n'
curl -d "@p.xml" -X POST http://localhost:1492/sifxml/PurchaseOrders
echo '\ncurl -X GET http://localhost:1492/sifxml/StaffPersonals/D3E34F41-9D75-101A-8C3D-00AA001A1652\n\n'
curl -X GET http://localhost:1492/sifxml/StaffPersonals/D3E34F41-9D75-101A-8C3D-00AA001A1652
echo '\ncurl -X GET http://localhost:1492/sifxml/PurchaseOrders\n\n'
curl -X GET http://localhost:1492/sifxml/PurchaseOrders
echo '\ncurl -X GET http://localhost:1492/sifxml/StaffPersonals\n\n'
curl -X GET http://localhost:1492/sifxml/StaffPersonals
echo '\ncp p.xml in\n\n'
cp p.xml 'in'

