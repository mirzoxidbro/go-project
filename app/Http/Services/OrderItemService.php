<?php

namespace App\Http\Services;

use App\Models\OrderItem;

class OrderItemService {

    public function index(){
        $orderitem = OrderItem::get();
        return $orderitem;
    }

    public function store(array $data){
        $orderitem = OrderItem::create($data);
        return $orderitem;
    }

    public function show($id){
        $orderitem = OrderItem::findOrFail($id);
        return $orderitem;
    }

    public function update(array $data, $id){
        $orderitem = OrderItem::findOrFail($id);
        $orderitem->update($data);
        return $orderitem;
    }

    public function destroy($id){
        $orderitem = OrderItem::findOrFail($id);
        $orderitem->delete();
    }
}