<?php

namespace App\Http\Services;

use App\Models\Product;

class ProductService {

    public function index(){
        $order = Product::where('actice', 'true');
        return $order;
    }

    public function store(array $data){
        $order = Product::create($data);
        return $order;
    }

    public function show($id){
        $order = Product::findOrFail($id);
        return $order;
    }

    public function update(array $data, $id){
        $order = Product::findOrFail($id);
        $order->update($data);
        return $order;
    }

    public function destroy($id){
        $order = Product::findOrFail($id);
        $order->delete();
    }
}