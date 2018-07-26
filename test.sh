echo '\ncurl -i -d "@purchaseorder.xml" -X POST http://localhost:1492/sifxml/Fred\n\n'
curl -i -d "@purchaseorder.xml" -X POST http://localhost:1492/sifxml/Fred
echo '\ncurl -i -d "@purchaseorder.xml" -X POST http://localhost:1492/sifxml/PurchaseOrders\n\n'
curl -i -d "@purchaseorder.xml" -X POST http://localhost:1492/sifxml/PurchaseOrders
echo '\ncurl -i -d "@staffpersonal.xml" -X POST http://localhost:1492/sifxml/PurchaseOrders\n\n'
curl -i -d "@staffpersonal.xml" -X POST http://localhost:1492/sifxml/PurchaseOrders
echo '\ncurl -i -d "@staffpersonal.xml" -X POST http://localhost:1492/sifxml/StaffPersonal\n\n'
curl -i -d "@staffpersonal.xml" -X POST http://localhost:1492/sifxml/StaffPersonals
echo '\ncurl -i -X GET http://localhost:1492/sifxml/StaffPersonals/D3E34F41-9D75-101A-8C3D-00AA001A1652\n\n'
curl -i -X GET http://localhost:1492/sifxml/StaffPersonals/D3E34F41-9D75-101A-8C3D-00AA001A1652
echo '\ncurl -i -d "@staffpersonal.xml" -H "mustUseAdvisory: true" -X POST http://localhost:1492/sifxml/StaffPersonals\n\n'
curl -i -d "@staffpersonal.xml" -H "mustUseAdvisory: true" -X POST http://localhost:1492/sifxml/StaffPersonals
echo '\ncurl -i -d "@staffpersonal.xml" -X POST http://localhost:1492/sifxml/StaffPersonals\n\n'
curl -i -d "@staffpersonal.xml" -X POST http://localhost:1492/sifxml/StaffPersonals
echo '\ncurl -i -d "@staffpersonals.xml" -X POST http://localhost:1492/sifxml/StaffPersonals\n\n'
curl -i -d "@staffpersonals.xml" -X POST http://localhost:1492/sifxml/StaffPersonals
echo '\ncurl -i -X GET http://localhost:1492/sifxml/PurchaseOrders\n\n'
curl -i -X GET http://localhost:1492/sifxml/PurchaseOrders
echo '\ncurl -i -X GET http://localhost:1492/sifxml/StaffPersonals\n\n'
curl -i -X GET http://localhost:1492/sifxml/StaffPersonals
echo '\ncurl -i -d "@purchaseorder.xml" -X PUT http://localhost:1492/sifxml/StaffPersonals/D3E34F41-9D75-101A-8C3D-00AA001A1652\n\n'
curl -i -d "@purchaseorder.xml" -X PUT http://localhost:1492/sifxml/StaffPersonals/D3E34F41-9D75-101A-8C3D-00AA001A1652
echo '\ncurl -i -d "@staffpersonal1.xml" -H "replacement: FULL" -X PUT http://localhost:1492/sifxml/StaffPersonals/D3E34F41-9D75-101A-8C3D-00AA001A1652\n\n'
curl -i -d "@staffpersonal1.xml" -H "replacement: FULL" -X PUT http://localhost:1492/sifxml/StaffPersonals/D3E34F41-9D75-101A-8C3D-00AA001A1652
echo '\ncurl -i -d "@staffpersonal2.xml" -X PUT http://localhost:1492/sifxml/StaffPersonals/D3E34F41-9D75-101A-8C3D-00AA001A1652\n\n'
curl -i -d "@staffpersonal2.xml" -X PUT http://localhost:1492/sifxml/StaffPersonals/D3E34F41-9D75-101A-8C3D-00AA001A1652
echo '\ncurl -i -d "@staffpersonal2.xml" -X DELETE http://localhost:1492/sifxml/StaffPersonals/D3E34F41-9D75-101A-8C3D-00AA001A1652\n\n'
curl -i -d "@staffpersonal2.xml" -X DELETE http://localhost:1492/sifxml/StaffPersonals/D3E34F41-9D75-101A-8C3D-00AA001A1652
echo '\ncurl -i -X GET http://localhost:1492/sifxml/StaffPersonals/D3E34F41-9D75-101A-8C3D-00AA001A1652\n\n'
curl -i -X GET http://localhost:1492/sifxml/StaffPersonals/D3E34F41-9D75-101A-8C3D-00AA001A1652
echo '\nrm in/purchaseorder.xml\n\n'
rm -f 'in/purchaseorder.xml'
echo '\ncp purchaseorder.xml in\n\n'
cp purchaseorder.xml 'in'

