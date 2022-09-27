<?php

namespace App\Http\Services;

use App\Models\Order;

class OrderService {

    public function index(){
        $order = Order::get();
        return $order;
    }

    public function store(array $data){
        $order = Order::create($data);
        return $order;
    }

    public function show($id){
        $order = Order::findOrFail($id);
        return $order;
    }

    public function update(array $data, $id){
        $order = Order::findOrFail($id);
        $order->update($data);
        return $order;
    }

    public function destroy($id){
        $order = Order::findOrFail($id);
        $order->delete();
    }
}