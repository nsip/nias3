#curl -d "@p.xml" -X POST http://localhost:1492/sifxml/Fred
echo '\ncurl -d "@p.xml" -X POST http://localhost:1492/sifxml/PurchaseOrder\n'
curl -d "@p.xml" -X POST http://localhost:1492/sifxml/PurchaseOrder
echo '\ncurl -X GET http://localhost:1492/sifxml/StaffPersonal/D3E34F41-9D75-101A-8C3D-00AA001A1652\n'
curl -X GET http://localhost:1492/sifxml/StaffPersonal/D3E34F41-9D75-101A-8C3D-00AA001A1652

