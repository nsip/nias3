#curl -d "@purchaseorder.xml" -X POST http://localhost:1492/sifxml/Fred
echo '\ncurl -d "@purchaseorder.xml" -X POST http://localhost:1492/sifxml/PurchaseOrders\n\n'
curl -d "@purchaseorder.xml" -X POST http://localhost:1492/sifxml/PurchaseOrders
# this object is inserted on startup
echo '\ncurl -X GET http://localhost:1492/sifxml/StaffPersonals/D3E34F41-9D75-101A-8C3D-00AA001A1652\n\n'
curl -X GET http://localhost:1492/sifxml/StaffPersonals/D3E34F41-9D75-101A-8C3D-00AA001A1652
echo '\ncurl -X GET http://localhost:1492/sifxml/PurchaseOrders\n\n'
curl -X GET http://localhost:1492/sifxml/PurchaseOrders
echo '\ncurl -X GET http://localhost:1492/sifxml/StaffPersonals\n\n'
curl -X GET http://localhost:1492/sifxml/StaffPersonals
echo '\ncurl -d "@purchaseorder.xml" -X PUT http://localhost:1492/sifxml/StaffPersonals/D3E34F41-9D75-101A-8C3D-00AA001A1652\n\n'
curl -d "@purchaseorder.xml" -X PUT http://localhost:1492/sifxml/StaffPersonals/D3E34F41-9D75-101A-8C3D-00AA001A1652
echo '\ncurl -d "@staffpersonal1.xml" -H "replacement: FULL" -X PUT http://localhost:1492/sifxml/StaffPersonals/D3E34F41-9D75-101A-8C3D-00AA001A1652\n\n'
curl -d "@staffpersonal1.xml" -H "replacement: FULL" -X PUT http://localhost:1492/sifxml/StaffPersonals/D3E34F41-9D75-101A-8C3D-00AA001A1652
echo '\ncurl -d "@staffpersonal2.xml" -X PUT http://localhost:1492/sifxml/StaffPersonals/D3E34F41-9D75-101A-8C3D-00AA001A1652\n\n'
curl -d "@staffpersonal2.xml" -X PUT http://localhost:1492/sifxml/StaffPersonals/D3E34F41-9D75-101A-8C3D-00AA001A1652
echo '\ncurl -d "@staffpersonal2.xml" -X DELETE http://localhost:1492/sifxml/StaffPersonals/D3E34F41-9D75-101A-8C3D-00AA001A1652\n\n'
curl -d "@staffpersonal2.xml" -X DELETE http://localhost:1492/sifxml/StaffPersonals/D3E34F41-9D75-101A-8C3D-00AA001A1652
echo '\ncurl -X GET http://localhost:1492/sifxml/StaffPersonals/D3E34F41-9D75-101A-8C3D-00AA001A1652\n\n'
curl -X GET http://localhost:1492/sifxml/StaffPersonals/D3E34F41-9D75-101A-8C3D-00AA001A1652
echo '\nrm in/purchaseorder.xml\n\n'
rm -f 'in/purchaseorder.xml'
echo '\ncp purchaseorder.xml in\n\n'
cp purchaseorder.xml 'in'

