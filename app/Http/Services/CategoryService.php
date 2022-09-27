<?php

namespace App\Http\Services;

use App\Models\Category;

class CategoryService {
    public function index(){
        $category = Category::where('active', 'true');
        return $category;
    }

    public function store(array $data){
        $category = Category::create($data);
        return $category;
    }

    public function show($id){
        $category = Category::findOrFail($id);
        return $category;
    }

    public function update(array $data, $id){
        $category = Category::findOrFail($id);
        $category->update($data);
        return $category;
    }

    public function destroy($id){
        $category = Category::findOrFail($id);
        $category->delete();
    }
}