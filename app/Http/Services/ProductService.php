<?php

namespace App\Http\Services;

use App\Models\Product;

class ProductService {

    public function index(){
        $product = Product::where('actice', 'true');
        return $product;
    }

    public function store(array $data){
        $product = Product::create($data);
        return $product;
    }

    public function show($id){
        $product = Product::findOrFail($id);
        return $product;
    }

    public function update(array $data, $id){
        $product = Product::findOrFail($id);
        $product->update($data);
        return $product;
    }

    public function destroy($id){
        $product = Product::findOrFail($id);
        $product->delete();
    }
}